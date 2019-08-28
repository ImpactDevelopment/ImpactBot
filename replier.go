package main

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	TIMEOUT = 30 * time.Second
	TRASH   = "🗑"

	general     = "208753003996512258"
	help        = "222120655594848256"
	bot         = "306182416329080833"
	donatorHelp = "583453983427788830"
	staff       = "308653317834145802" // for testing
)

var channels = []string{general, help, bot, donatorHelp, staff}

// a map from ID of a message I sent, to the ID of who is allowed to delete it (aka who sent the message that I was responding to)
var messageSender = make(map[string]string)
var messageSenderLock sync.Mutex

func onMessageReactedTo(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	messageSenderLock.Lock()
	defer messageSenderLock.Unlock()
	origAuthor, ok := messageSender[m.MessageID]
	if !ok {
		return // this wasn't us
	}
	if m.Emoji.Name == TRASH && (isSupport(m.UserID) || m.UserID == origAuthor) && m.UserID != myselfID {
		go discord.ChannelMessageDelete(m.ChannelID, m.MessageID) // sometimes errors since it was already trashcanned, dont spam logs with this error its too common
	}
}

func onMessageSent(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault {
		return // wtf
	}
	author := msg.Author.ID

	// Don't talk to oneself
	if author == myselfID {
		return
	}

	// Unless we're being spoken to
	if !triggeredManually(msg) {
		// Don't talk where we're not welcome
		if !includes(channels, msg.ChannelID) {
			return
		}

		// Ignore messages from ‘know-it-all’s
		if isSupport(author) {
			return
		}
	}

	// Phew, actually start doing stuff
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
		return // if this failed, msg will be nil, so we cannot continue!
	}

	// Add a trashcan icon if the message wasn't triggered manually
	// Keep track of who is allowed to delete the message too
	if !triggeredManually(msg) {
		messageSender[msg.ID] = author
		err = discord.MessageReactionAdd(msg.ChannelID, msg.ID, TRASH)
		if err != nil {
			log.Println(err)
		}

		// Delete the entry from our messageSender map after the TIMEOUT
		go func() {
			time.Sleep(TIMEOUT)
			messageSenderLock.Lock()
			defer messageSenderLock.Unlock()
			delete(messageSender, msg.ID)
		}()
	}
}

func triggeredManually(msg *discordgo.Message) bool {
	// TODO other methods of manual triggering, e.g. i!commands
	return mentionsMe(msg)
}
