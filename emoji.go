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
		err := discord.MessageDelete(m.GuildID, m.Author.ID, "Sending a non emoji message in the emoji-only channel")
		if err != nil {
			return
		}
	}
}


func onMessageEdited(session *discordgo.Session, m *discordgo.MessageUpdate) {
	emojiMsg(m.Message)
}

