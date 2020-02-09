package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Censorship struct {
	name        string
	bannedWords []string
}

var censor = map[string]Censorship{
	"563138570953687061": Censorship{"Bella", []string{"kami", "blue", "力ミ", "ブル"}},
	"209785549010108416": Censorship{"Arisa", []string{"loli", "smh"}},
	"207833493651193856": Censorship{"XPHonza", []string{"boomer"}},
	"297773911158816769": Censorship{"leijurv", []string{"not allowed to say this"}},
}

func onMessageSent3(session *discordgo.Session, m *discordgo.MessageCreate) {
	enforcement(session, m.Message)
}

func onMessageUpdate(session *discordgo.Session, m *discordgo.MessageUpdate) {
	enforcement(session, m.Message)
}

func enforcement(session *discordgo.Session, msg *discordgo.Message) {
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault {
		return // wtf
	}

	censorship, ok := censor[msg.Author.ID]
	if !ok {
		return
	}

	for _, bannedWord := range censorship.bannedWords {
		if strings.Contains(strings.ToLower(msg.Content), strings.ToLower(bannedWord)) {
			session.ChannelMessageDelete(msg.ChannelID, msg.ID)
			resp(msg.ChannelID, "Note: a message containing \""+bannedWord+"\" from "+censorship.name+" was deleted")
			return
		}
	}
}
