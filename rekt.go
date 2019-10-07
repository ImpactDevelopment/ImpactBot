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
	var ban bool
	if content[:5] == "kick " {
		ban = false
	} else {
		if content[:4] == "ban " {
			ban = true
		} else {
			return
		}
	}

	author, err := GetMember(msg.Author.ID)
	if err != nil {
		return
	}
	if !isStaff(author) {
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
	if !evalRatelimit(msg.Author.ID) {
		resp(msg.ChannelID, "Too soon")
		return
	}

	providedReason := strings.TrimSpace(content[strings.Index(content, ">")+1:])

	providedReason = "by @" + msg.Author.Username + "#" + msg.Author.Discriminator + " for reason: " + providedReason

	text := "@" + user.Username + "#" + user.Discriminator + " was "
	if ban {
		text += "banned"
	} else {
		text += "kicked"
	}
	text += " " + providedReason

	if ban {
		err = discord.GuildBanCreateWithReason(m.GuildID, user.ID, providedReason, 0)
	} else {
		err = discord.GuildMemberDeleteWithReason(m.GuildID, user.ID, providedReason)
	}

	if err != nil {
		resp(msg.ChannelID, "ERROR "+err.Error())
		return
	}
	resp(msg.ChannelID, text)
}
