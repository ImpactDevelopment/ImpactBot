package main

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

const (
	RATELIMIT         = 5 * time.Minute
	TEMPMUTE_DURATION = 3 * time.Hour
	UNMUTE_INTERVAL   = 10 * time.Second
)

var ratelimit = make(map[string]int64)
var ratelimitLock sync.Mutex

var mentionRegex = regexp.MustCompile(`^<(?P<Type>[#@])!?(?P<ID>\d+)>$`)

func evalRatelimit(author string) bool {
	ratelimitLock.Lock()
	defer ratelimitLock.Unlock()

	until := ratelimit[author]
	if until < time.Now().UnixNano() { // defaults to 0 so this works properly
		ratelimit[author] = time.Now().Add(RATELIMIT).UnixNano()
		return true
	}

	return false
}

func resp(ch string, text string) error {
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       prettyembedcolor,
		Description: text,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "♿ Impact Client ♿",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
	_, err := discord.ChannelMessageSendEmbed(ch, embed)
	return err
}

// Turns the first one or two args into users and/or channels and also returns whatever args weren't consumed
func getUserAndChannelAndArgs(args []string) (user *discordgo.User, channel *discordgo.Channel, remainingArgs []string) {
	remainingArgs = args
	if len(args) < 1 {
		return
	}
	user, channel = getUserOrChannelForArg(args[0])
	if user != nil || channel != nil {
		// Consume an arg
		remainingArgs = remainingArgs[1:]
		// Don't return since we want to try the second arg too
	} else {
		// No match on first arg so don't try to match second arg
		return
	}

	// getUserOrChannelForArg always has one nil arg, so if-else instead of if-elseif is fine
	if len(args) < 2 {
		return
	}
	if user == nil {
		user, _ = getUserOrChannelForArg(args[1])
		if user != nil {
			// Consume an arg
			remainingArgs = remainingArgs[1:]
		}
	} else {
		_, channel = getUserOrChannelForArg(args[1])
		if channel != nil {
			// Consume an arg
			remainingArgs = remainingArgs[1:]
		}
	}
	return
}

// Send a blocking api request if a match is found
func getUserOrChannelForArg(arg string) (*discordgo.User, *discordgo.Channel) {
	match := findNamedMatches(mentionRegex, arg)
	if match["ID"] == "" {
		return nil, nil
	}
	switch match["Type"] {
	// Sends a blocking API request
	case "#":
		{
			channel, err := discord.Channel(match["ID"])
			if err == nil {
				return nil, channel
			}
		}
	case "@":
		{
			user, err := discord.User(match["ID"])
			if err == nil {
				return user, nil
			}
		}
	}

	return nil, nil
}

// Gets the mute role associated with the given channel
// Returns an error if no matching role exists
func getMuteRoleForChannel(channel *discordgo.Channel) (role string, err error) {
	if channel == nil {
		role = muteRoles[""]
		return
	}
	var ok bool // Avoid shadowing role with :=
	if role, ok = muteRoles[channel.ID]; !ok {
		err = fmt.Errorf("unable to find mute role for channel %s", channel.Mention())
	}
	return
}

func muteHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	user, channel, remainingArgs := getUserAndChannelAndArgs(args[1:])
	if user == nil {
		return errors.New("First argument should mention user or channel")
	}
	if len(remainingArgs) < 1 {
		return errors.New("Give a reason")
	}

	muteRole, err := getMuteRoleForChannel(channel)
	if err != nil {
		return fmt.Errorf("Can't mute from %s yet", channel.Mention())
	}

	// Reasons are important
	var reason strings.Builder
	reason.WriteString(args[0])
	reason.WriteString(" has been issued to " + user.Username)
	if channel != nil {
		reason.WriteString(" from channel " + channel.Mention())
	}
	reason.WriteString(" by @" + msg.Author.Username + "#" + msg.Author.Discriminator)
	reason.WriteString(" for reason: " + strings.Join(remainingArgs, " "))
	providedReason := reason.String()

	// Direct message the user being muted
	DM, err := discord.UserChannelCreate(user.ID) // only creates it if it doesn"t already exist
	if err == nil {
		// if there is an error DMing them, we still want to ban them, they just won't know why
		err = resp(DM.ID, providedReason)
		if err != nil {
			fmt.Printf("Error direct messaging %s#%s: %s\n", user.Username, user.Discriminator, err.Error())
		}
	}

	// Support can tempmute, but only on users without roles
	if strings.ToLower(args[0]) == "tempmute" {
		member, err := GetMember(user.ID)
		if err != nil {
			return err
		}
		if IsUserLowerThan(caller, Moderator) && len(member.Roles) > 0 {
			return errors.New("They have role(s)")
		}
		if IsUserLowerThan(caller, Moderator) && !evalRatelimit(msg.Author.ID) {
			return errors.New("Too soon")
		}

		if DB == nil {
			return errors.New("I have no database, so I cannot tempmute")
		}

		// Row values
		var (
			id         uuid.UUID
			userId     = user.ID
			channelId  sql.NullString
			expiration = time.Now().Add(TEMPMUTE_DURATION)
		)
		if channel != nil {
			channelId = sql.NullString{
				String: channel.ID,
				Valid:  true,
			}
		}

		// Check if we have a matching mute already
		err = DB.QueryRow("SELECT id from tempmutes WHERE discord_id=$1 AND channel_id=$2", userId, channelId).Scan(&id)

		if err == nil {
			// Update existing entry
			_, err = DB.Exec("UPDATE tempmutes SET expiration=$2 WHERE id=$1", id, expiration)
			if err != nil {
				return err
			}
		} else if errors.Is(err, sql.ErrNoRows) {
			// Insert new entry
			_, err = DB.Exec("INSERT INTO tempmutes (discord_id, channel_id, expiration) VALUES ($1, $2, $3)", userId, channelId, expiration)
			if err != nil {
				return err
			}
		} else {
			// Unknown error
			return err
		}
	}
	err = discord.GuildMemberRoleAdd(msg.GuildID, user.ID, muteRole)
	if err != nil {
		return err
	}

	_ = resp(FORWARD_TO, providedReason)

	_ = resp(msg.ChannelID, providedReason)
	return nil
}

func unmuteHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) (err error) {
	var unmuteAll bool
	user, channel, remainingArgs := getUserAndChannelAndArgs(args[1:])
	if user == nil {
		return errors.New("First argument should mention user")
	}
	if channel == nil && len(remainingArgs) == 1 && "all" == strings.ToLower(remainingArgs[0]) {
		unmuteAll = true
	} else if len(remainingArgs) > 0 {
		return fmt.Errorf("Unexpected arguments \"%s\"", strings.Join(remainingArgs, " "))
	}

	var member *discordgo.Member
	member, err = GetMember(user.ID)
	if err != nil {
		return
	}

	// produce a list of muted channels for the command output
	var fullMute bool
	var channels []string

	if unmuteAll {
		// Remove all muted roles
		for mutedChannel, muteRole := range muteRoles {
			if hasRole(member, Role{ID: muteRole}) {
				if mutedChannel == "" {
					fullMute = true
				} else {
					channels = append(channels, mutedChannel)
				}
				err = discord.GuildMemberRoleRemove(msg.GuildID, user.ID, muteRole)
			}
		}
	} else {
		// unmute specified channel (or server-wide for nil)
		muteRole, err := getMuteRoleForChannel(channel)
		if err != nil {
			if channel != nil {
				return fmt.Errorf("Can't unmute from %s yet", channel.Mention())
			} else {
				return err // Unknown error
			}
		}

		// Check the user is actually muted...
		if !hasRole(member, Role{ID: muteRole}) {
			if channel == nil {
				return fmt.Errorf("%s isn't muted serverwide", user.Username)
			} else {
				return fmt.Errorf("%s isn't muted in %s", user.Username, channel.Mention())
			}
		}

		// Log the unmute
		if channel == nil {
			fullMute = true
		} else {
			channels = append(channels, channel.ID)
		}

		// Unmute them
		err = discord.GuildMemberRoleRemove(msg.GuildID, user.ID, muteRole)
	}

	// Construct a reply out of the unmuted channels slice
	var reply strings.Builder
	reply.WriteString("User " + user.Username + " has been ")
	if fullMute {
		reply.WriteString("unmuted serverwide")
		if len(channels) > 0 {
			reply.WriteString(" and ")
		}
	}
	if len(channels) > 0 {
		reply.WriteString("unmuted from ")
		for index, id := range channels {
			if index > 0 {
				reply.WriteString(", ")
				if index == len(channels)-1 {
					reply.WriteString("& ")
				}
			}
			reply.WriteString(fmt.Sprintf("<#%s>", id))
		}
	}
	reply.WriteString(" by @" + msg.Author.Username + "#" + msg.Author.Discriminator)
	reply.WriteString("\n")

	// Direct message the user being unmuted
	DM, err := discord.UserChannelCreate(user.ID)
	if err == nil {
		err = resp(DM.ID, reply.String())
		if err != nil {
			s := fmt.Sprintf("Error direct messaging %s#%s: %s", user.Username, user.Discriminator, err.Error())
			fmt.Println(s)
			reply.WriteString(s + "\n")
		}
	}

	// Respond in chat & #impactbot-log
	_ = resp(FORWARD_TO, reply.String())
	_ = resp(msg.ChannelID, reply.String())

	return nil
}

// tbh should this be separate handlers??
// or maybe multiple handlers here is stupid, this is mostly a copy of mute handler :\
func rektHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	user, channel, remainingArgs := getUserAndChannelAndArgs(args[1:])
	if user == nil {
		return errors.New("First argument should mention user")
	}
	if channel != nil {
		return errors.New(args[0] + " does not support channel mentions")
	}
	if len(remainingArgs) < 1 {
		return errors.New("Give a reason")
	}

	// Reasons are important
	providedReason := args[0] + " has been issued to " + user.Username + " by @" + msg.Author.Username + "#" + msg.Author.Discriminator + " for reason: " + strings.Join(remainingArgs, " ")

	// Direct message the user being rekt
	DM, err := discord.UserChannelCreate(user.ID) // only creates it if it doesn"t already exist
	if err == nil {
		// if there is an error DMing them, we still want to ban them, they just won't know why
		err = resp(DM.ID, providedReason)
		if err != nil {
			fmt.Printf("Error direct messaging %s#%s: %s\n", user.Username, user.Discriminator, err.Error())
		}
	}

	switch args[0] {
	case "ban":
		err = discord.GuildBanCreateWithReason(msg.GuildID, user.ID, providedReason, 0)
	case "kick":
		err = discord.GuildMemberDeleteWithReason(msg.GuildID, user.ID, providedReason)
	}

	if err != nil {
		return err
	}

	_ = resp(FORWARD_TO, providedReason)

	_ = resp(msg.ChannelID, providedReason)
	return nil
}

// unmuteCallback is called every UNMUTE_INTERVAL to unmute any expired temp mutes
func unmuteCallback() {
	if DB == nil {
		return
	}

	// Get all expired rows
	now := time.Now()
	rows, err := DB.Query("SELECT id, discord_id, channel_id FROM tempmutes WHERE expiration < $1", now)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Error querying expired tempmutes", err)
		}
		return
	}

	for rows.Next() {
		var (
			id        uuid.UUID
			discordId string
			channelId sql.NullString
		)

		err := rows.Scan(&id, &discordId, &channelId)
		if err != nil {
			fmt.Println("Error scanning tempmute entry", err)
			continue
		}

		// We're handling the user, so delete from db.
		// If we fail to unmute, at least we won't end up handling them forever
		_, err = DB.Exec("DELETE FROM tempmutes WHERE id = $1", id)
		if err != nil {
			fmt.Println("Error deleting tempmute entry", err)
			continue
		}

		// Get channel and mute role
		var channel *discordgo.Channel
		if channelId.Valid {
			channel, err = discord.Channel(channelId.String)
			if err != nil {
				fmt.Println("Error getting tempmute channel from channel id "+channelId.String, err)
				continue
			}
		}
		muteRole, err := getMuteRoleForChannel(channel)
		if err != nil {
			fmt.Println(fmt.Sprintf("Invalid mute role for channel id \"%s\"\n", channelId.String), err)
			continue
		}

		// Construct a message to the unmuted user
		var message strings.Builder
		message.WriteString(fmt.Sprintf("Your temp mute"))
		if channel != nil {
			message.WriteString(fmt.Sprintf(" from %s", channel.Mention()))
		}
		message.WriteString(" has expired!\n")

		{ // Log the unmute
			var username = "user"
			if user, _ := discord.User(discordId); user != nil {
				username = fmt.Sprintf("@%s#%s", user.Username, user.Discriminator)
			}
			if channel == nil {
				fmt.Printf("Processing unmute for %s (%s) serverwide\n", username, discordId)
			} else {
				fmt.Printf("Processing unmute for %s (%s) from channel #%s (%s)\n", username, discordId, channel.Name, channel.ID)
			}
		}

		// Do the unmute
		err = discord.GuildMemberRoleRemove(IMPACT_SERVER, discordId, muteRole)
		if err != nil {
			fmt.Println("Could not remove mute role \""+muteRole+"\" from user \""+discordId+"\"", err)
			message.WriteString("But the bot failed to unmute you! Please show this message to a moderator.\n")
		}

		// DM message to user
		dm, err := discord.UserChannelCreate(discordId)
		if err != nil {
			continue // guess we can't let em know
		}
		_ = resp(dm.ID, message.String())
	}
}

func init() {
	if DB == nil {
		fmt.Println("WARNING: No DB when initialising tempmutes callback, either rekt.go was initialised before db.go or there is no DB")
	}
	go func() {
		ticker := time.NewTicker(UNMUTE_INTERVAL)
		for range ticker.C {
			unmuteCallback()
		}
	}()
}
