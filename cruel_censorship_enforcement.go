package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Censorship struct {
	name string
	bannedWords []string
}

var censor = map[string]Censorship{
	"563138570953687061": Censorship{"Bella", []string{"kami", "blue"}},
	"209785549010108416": Censorship{"Arisa", []string{"lewd", "loli", "𝗹𝗲𝘄𝗱", "𝗹𝗼𝗹𝗶"}},
}

func onMessageSent3(session *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
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
			resp(msg.ChannelID, "Note: a message containing \"" + bannedWord + "\" from " + censorship.name + " was deleted")
			return
		}
	}
}
