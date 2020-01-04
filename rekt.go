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

func onMessageSent3(session *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault || msg.Author.ID == myselfID || m.GuildID != IMPACT_SERVER {
		return // wtf
	}

	content := msg.Content

	if strings.HasPrefix(content, "i!") { // bot woke
		fields := strings.Fields(content[2:])
		command := strings.ToLower(fields[0])
		if contains([]string{"kick", "ban", "tempmute", "mute", "unmute"}, command) { // we don't want role checking outside of these commands
			author, err := GetMember(msg.Author.ID)
			if err != nil {
				return
			}
			// Allow support to run `tempmute` and mods to run anything
			if !(command == "tempmute" && IsUserAtLeast(author, Support)) || IsUserAtLeast(author, Moderator) {
				return
			}
			if len(msg.Mentions) != 1 {
				resp(msg.ChannelID, "Mention exactly one user")
				return
			}
			user := msg.Mentions[0]
			member, err := GetMember(user.ID)
			if err != nil {
				return
			}
			if !IsUserAtLeast(author, Moderator) && len(member.Roles) > 0 && command != "unmute" {
				resp(msg.ChannelID, "They have role(s)")
				return
			}
			if !IsUserAtLeast(author, Moderator) && !evalRatelimit(msg.Author.ID) && command != "unmute" {
				resp(msg.ChannelID, "Too soon")
				return
			}
			providedReason := strings.TrimSpace(content[strings.Index(content, ">")+1:])
			if providedReason == "" {
				resp(msg.ChannelID, "Give a reason")
				return
			}
			providedReason = command + " has been issued to " + user.Username + " by @" + msg.Author.Username + "#" + msg.Author.Discriminator + " for reason: " + providedReason

			DM, err := discord.UserChannelCreate(user.ID) // only creates it if it doesn"t already exist
			if err == nil {
				// if there is an error DMing them, we still want to ban them, they just won't know why
				resp(DM.ID, providedReason)
			}

			switch command {
			case "ban":
				err = discord.GuildBanCreateWithReason(m.GuildID, user.ID, providedReason, 0)
			case "kick":
				err = discord.GuildMemberDeleteWithReason(m.GuildID, user.ID, providedReason)
			case "tempmute":
				if DB == nil {
					err = errors.New("I have no database, so I cannot tempmute")
					break
				}
				_, err = DB.Exec("INSERT INTO tempmutes(discord_id, expiration) VALUES ($1, $2) ON CONFLICT(discord_id) DO UPDATE SET expiration = EXCLUDED.expiration", user.ID, time.Now().Unix()+6*3600)
				if err != nil {
					break
				}
				fallthrough
			case "mute":
				err = discord.GuildMemberRoleAdd(m.GuildID, user.ID, MUTE_ROLE)
			case "unmute":
				err = discord.GuildMemberRoleRemove(m.GuildID, user.ID, MUTE_ROLE)
			}
			if err != nil {
				resp(msg.ChannelID, "ERROR "+err.Error())
				return
			}

			resp(FORWARD_TO, providedReason)

			resp(msg.ChannelID, providedReason)
		}
	}
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
