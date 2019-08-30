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
	testing     = "617066818925756506"
)

var channels = []string{general, help, bot, donatorHelp, testing}

// a map from ID of a message I sent, to the ID of who is allowed to delete it (aka who sent the message that I was responding to)
var messageSender = make(map[string]string)
var messageSenderLock sync.Mutex

func onMessageReactedTo(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	messageSenderLock.Lock()
	defer messageSenderLock.Unlock()

	// If the reaction isn't trash we don't care
	if reaction.Emoji.Name != TRASH {
		return
	}

	// Get the author of the trigger message
	author, ok := messageSender[reaction.MessageID]
	if !ok {
		return // this wasn't us
	}

	// Get the reply we sent
	reply, err := session.ChannelMessage(reaction.ChannelID, reaction.MessageID)
	if err != nil {
		return //wtf
	}

	// If we didn't send the reply or we added the reaction
	if reply.Author.ID != myselfID || reaction.UserID == myselfID {
		return
	}

	// Filter approved users
	if reaction.UserID != author && !isSupport(reaction.UserID) {
		return
	}

	// Delete the reply
	// sometimes errors since it was already trashcanned, dont spam logs with this error its too common
	go session.ChannelMessageDelete(reply.ChannelID, reply.ID)
}

func onMessageSent(session *discordgo.Session, m *discordgo.MessageCreate) {
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
	reply, err := session.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		log.Println(err)
		return // if this failed, msg will be nil, so we cannot continue!
	}

	// Add a trashcan icon if the message wasn't triggered manually
	// Keep track of who is allowed to delete the message too
	if !triggeredManually(msg) {
		err = session.MessageReactionAdd(reply.ChannelID, reply.ID, TRASH)
		if err != nil {
			log.Println(err)
		}

		// Add the message to the sender map then delete it later
		messageSender[reply.ID] = author
		go func() {
			time.Sleep(TIMEOUT)
			messageSenderLock.Lock()
			defer messageSenderLock.Unlock()
			delete(messageSender, reply.ID)
		}()
	}
}

func triggeredManually(msg *discordgo.Message) bool {
	// TODO other methods of manual triggering, e.g. i!commands
	return mentionsMe(msg)
}
