package main

import (
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	TIMEOUT = 20 * time.Second
	TRASH   = "🗑"

	general     = "208753003996512258"
	help        = "222120655594848256"
	bot         = "306182416329080833"
	donatorHelp = "583453983427788830"
)

var channels = []string{general, help, bot, donatorHelp}

func onMessageReactedTo(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.Emoji.Name == TRASH && isSupport(m.UserID) && m.UserID != myselfID {
		discord.ChannelMessageDelete(m.ChannelID, m.MessageID) // sometimes errors since it was already trashcanned, dont spam logs with this error its too common
	}
}

func onMessageSent(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil {
		log.Println("wtf")
		return
	}
	if !includes(channels, msg.ChannelID) || msg.Type != discordgo.MessageTypeDefault || msg.Author == nil || msg.Author.ID == myselfID {
		return
	}
	if isSupport(msg.Author.ID) && !mentionsMe(msg) {
		return
	}
	response := ""
	for _, reply := range replies {
		if includes(reply.exclude, msg.ChannelID) {
			continue
		}
		if reply.regex.MatchString(strings.ToLower(msg.Content)) {
			response += reply.message + " "
		}
	}
	if response == "" {
		return
	}
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       prettyembedcolor,
		Description: response,
	}
	msg, err := discord.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		log.Println(err)
	}
	err = discord.MessageReactionAdd(msg.ChannelID, msg.ID, TRASH)
	if err != nil {
		log.Println(err)
	}
	go func() {
		time.Sleep(TIMEOUT)
		err := discord.ChannelMessageDelete(msg.ChannelID, msg.ID)
		if err != nil {
			log.Println(err)
		}
	}()
}
