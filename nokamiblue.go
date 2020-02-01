package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

const BELLA = "563138570953687061"


func onMessageSent3(session *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault {
		return // wtf
	}

	if msg.Author.ID != BELLA {
		return
	}

	if strings.Contains(strings.ToLower(msg.Content), "kami") || strings.Contains(strings.ToLower(msg.Content), "blue") {
		session.ChannelMessageDelete(msg.ChannelID, msg.ID)
		resp(msg.ChannelID, "Note: a message containing \"kami\" or \"blue\" from Bella was deleted")
	}
}
