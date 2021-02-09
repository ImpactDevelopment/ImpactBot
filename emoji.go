package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
	"regexp"
	"strings"
	"unicode"
)

var discordEmote = regexp.MustCompile(`<a?:[a-zA-Z0-9_]+:\d+>`)
func emojiMsg(m *discordgo.Message){
	if m.ChannelID != "808248247520985130" {
		return
	}

	stripped := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, discordEmote.ReplaceAllString(emoji.RemoveAll(m.Content), ""))

	log.Println("Original message:", m.Content)
	log.Println("Stripped:", stripped)

	if stripped != "" {
		err := discord.GuildMemberDeleteWithReason(m.GuildID, m.Author.ID, "Sending a non emoji message in the emoji-only channel")
		if err != nil {
			resp(m.ChannelID, "ERROR "+err.Error())
			return
		}
		resp(m.ChannelID, "User <@" + m.Author.ID + "> was kicked for sending https://discord.com/channels/208753003996512258/808248247520985130/"+m.ID+" because it has non emoji characters: `" + stripped + "`")
	}
}


func onMessageEdited(session *discordgo.Session, m *discordgo.MessageUpdate) {
	emojiMsg(m.Message)
}

