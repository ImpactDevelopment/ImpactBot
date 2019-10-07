package main

import (
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	RATELIMIT = 5 * time.Minute
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
	}
	discord.ChannelMessageSendEmbed(ch, embed)
}

func onMessageSent3(session *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault || msg.Author.ID == myselfID || m.GuildID != IMPACT_SERVER {
		return // wtf
	}

	content := msg.Content
	if len(content) < 5 {
		return
	}
	if content[:2] == "i!" { // bot woke
		var fields = strings.Fields(content)
		var command = fields[0][2:]
		if command == "kick" || command == "ban" || command == "mute" || command == "unmute" { // we don't want role checking outside of these commands
			author, err := GetMember(msg.Author.ID)
			if err != nil || !isStaff(author) {
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
			if len(member.Roles) > 0 {
				resp(msg.ChannelID, "They have role(s)")
				return
			}
			if !hasRole(author, STAFF["moderator"], STAFF["developer"]) && !evalRatelimit(msg.Author.ID) {
				resp(msg.ChannelID, "Too soon")
				return
			}
			providedReason := strings.TrimSpace(content[strings.Index(content, ">")+1:])
			providedReason = command + " has been issued to " + user.Username + " by @" + msg.Author.Username + "#" + msg.Author.Discriminator + " for reason : " + providedReason
			switch command {
			case "ban":
				err = discord.GuildBanCreateWithReason(m.GuildID, user.ID, providedReason, 0)
				break
			case "kick":
				err = discord.GuildMemberDeleteWithReason(m.GuildID, user.ID, providedReason)
				break
			case "mute":
				err = discord.GuildMemberRoleAdd(m.GuildID, user.ID, "dummyRoleID")
				break
			case "unmute":
				err = discord.GuildMemberRoleRemove(m.GuildID, user.ID, "dummyRoleID")
				break
			}
			if err != nil {
				resp(msg.ChannelID, "ERROR "+err.Error())
				return
			}

			resp(msg.ChannelID, providedReason)
		}
	}
}
