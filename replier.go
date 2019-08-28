package main

import (
	"log"
	"strings"
	"sync"
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

// a map from ID of a message I sent, to the ID of who is allowed to delete it (aka who sent the message that I was responding to)
var messageSender = make(map[string]string)
var messageSenderLock sync.Mutex

func onMessageReactedTo(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	messageSenderLock.Lock()
	defer messageSenderLock.Unlock()
	origAuthor, ok := messageSender[m.UserID]
	if !ok {
		return // this wasn't us
	}
	if m.Emoji.Name == TRASH && (isSupport(m.UserID) || m.UserID == origAuthor) && m.UserID != myselfID {
		discord.ChannelMessageDelete(m.ChannelID, m.MessageID) // sometimes errors since it was already trashcanned, dont spam logs with this error its too common
	}
}

func onMessageSent(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil {
		log.Println("wtf")
		return
	}
	author := msg.Author.ID
	if !includes(channels, msg.ChannelID) || msg.Type != discordgo.MessageTypeDefault || msg.Author == nil || author == myselfID {
		return
	}
	mentionedMe := mentionsMe(msg)
	if isSupport(author) && !mentionedMe {
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
	messageSenderLock.Lock()
	defer messageSenderLock.Unlock()
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       prettyembedcolor,
		Description: response,
	}
	msg, err := discord.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		log.Println(err)
	}
	messageSender[msg.ID] = author
	if !mentionedMe {
		err = discord.MessageReactionAdd(msg.ChannelID, msg.ID, TRASH)
		if err != nil {
			log.Println(err)
		}
	}
	go func() {
		time.Sleep(TIMEOUT)
		if !mentionedMe {
			err := discord.ChannelMessageDelete(msg.ChannelID, msg.ID)
			if err != nil {
				log.Println(err)
			}
		}
		messageSenderLock.Lock()
		defer messageSenderLock.Unlock()
		delete(messageSender, msg.ID)
	}()
}
