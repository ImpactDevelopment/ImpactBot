package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	RATELIMIT = 5 * time.Minute
	MUTE_ROLE = "630800201015361566"
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

// TODO add channel specific mute mode
func muteHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	user, channel, remainingArgs := getUserAndChannelAndArgs(args[1:])
	if user == nil {
		return errors.New("First argument should mention user")
	}
	if channel != nil { //TODO
		return errors.New("Channel mentions not supported... for now")
	}
	if len(remainingArgs) < 1 {
		return errors.New("Give a reason")
	}

	// Reasons are important
	providedReason := args[0] + " has been issued to " + user.Username + " by @" + msg.Author.Username + "#" + msg.Author.Discriminator + " for reason: " + strings.Join(remainingArgs, " ")

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
		_, err = DB.Exec("INSERT INTO tempmutes(discord_id, expiration) VALUES ($1, $2) ON CONFLICT(discord_id) DO UPDATE SET expiration = EXCLUDED.expiration", user.ID, time.Now().Add(3*time.Hour).Unix())
		if err != nil {
			return err
		}
	}
	err = discord.GuildMemberRoleAdd(msg.GuildID, user.ID, MUTE_ROLE)
	if err != nil {
		return err
	}

	_ = resp(FORWARD_TO, providedReason)

	_ = resp(msg.ChannelID, providedReason)
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
	case "unmute":
		err = discord.GuildMemberRoleRemove(msg.GuildID, user.ID, MUTE_ROLE)
	}

	if err != nil {
		return err
	}

	_ = resp(FORWARD_TO, providedReason)

	_ = resp(msg.ChannelID, providedReason)
	return nil
}

func init() {
	if DB == nil {
		fmt.Println("Tempmutes will never end since I don't have access to a database lol")
		return
	}
	go func() {
		for {
			time.Sleep(10 * time.Second)
			now := time.Now().Unix()
			var id string
			err := DB.QueryRow("SELECT discord_id FROM tempmutes WHERE expiration < $1", now).Scan(&id)
			if err != nil {
				continue // probably sql.ErrNoRows
			}
			_, err = DB.Exec("DELETE FROM tempmutes WHERE discord_id = $1", id)
			if err != nil {
				fmt.Println("Couldn't delete?", err)
				continue
			}
			fmt.Println("Processing temp unmute for", id)
			err = discord.GuildMemberRoleRemove(IMPACT_SERVER, id, MUTE_ROLE)
			if err != nil {
				fmt.Println("Could not remove mute role", err)
				continue
			}
			DM, err := discord.UserChannelCreate(id) // only creates it if it doesn"t already exist
			if err != nil {
				// guess we can't let em know
				continue
			}
			_ = resp(DM.ID, "Your temorary mute is over")
		}
	}()
}
