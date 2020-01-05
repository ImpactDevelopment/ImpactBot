package main

import (
	"errors"
	"fmt"
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

func resp(ch string, text string) {
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       prettyembedcolor,
		Description: text,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "♿ Impact Client ♿",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
	discord.ChannelMessageSendEmbed(ch, embed)
}

func contains(list []string, it string) bool {
	for _, s := range list {
		if s == it {
			return true
		}
	}
	return false
}

// TODO get without using Message.Mentions?
func getUserFromArgs(msg *discordgo.Message, args []string) (*discordgo.User, error) {
	if len(msg.Mentions) != 1 {
		return nil, errors.New("Mention exactly one user")
	}
	user := msg.Mentions[0]
	if args[1] != fmt.Sprintf("<@!%s>", user.ID) {
		return nil, errors.New("First argument should mention user")
	}
	return user, nil
}

// TODO get without using Message.Mentions?
func getReasonFromArgs(args []string) string {
	index := 1
	// Find the first index that isn't a user/channel mention
	for index < len(args) && strings.HasPrefix(args[index], "<") && strings.HasSuffix(args[index], ">") {
		index++
	}
	if index >= len(args) {
		return ""
	}
	return strings.Join(args[index:], " ")
}

// TODO add channel specific mute mode
func muteHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	user, err := getUserFromArgs(msg, args)
	if err != nil {
		return err
	}

	// Reasons are important
	reason := getReasonFromArgs(args)
	if reason == "" {
		return errors.New("Give a reason")
	}
	providedReason := args[0] + " has been issued to " + user.Username + " by @" + msg.Author.Username + "#" + msg.Author.Discriminator + " for reason: " + reason

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

		DM, err := discord.UserChannelCreate(user.ID) // only creates it if it doesn"t already exist
		if err == nil {
			// if there is an error DMing them, we still want to ban them, they just won't know why
			resp(DM.ID, providedReason)
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

	resp(FORWARD_TO, providedReason)

	resp(msg.ChannelID, providedReason)
	return nil
}

// tbh should this be separate handlers??
func rektHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	user, err := getUserFromArgs(msg, args)
	if err != nil {
		return err
	}

	// Reasons are important
	reason := getReasonFromArgs(args)
	if reason == "" {
		return errors.New("Give a reason")
	}
	providedReason := args[0] + " has been issued to " + user.Username + " by @" + msg.Author.Username + "#" + msg.Author.Discriminator + " for reason: " + reason

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

	resp(FORWARD_TO, providedReason)

	resp(msg.ChannelID, providedReason)
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
			resp(DM.ID, "Your temorary mute is over")
		}
	}()
}
