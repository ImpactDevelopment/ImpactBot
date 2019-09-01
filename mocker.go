package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	FORWARD_TO = "308653317834145802"
)

// inform when someone DMs the bot because the messages are humorous
func onMessageSent2(session *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault {
		return // wtf
	}
	author := msg.Author.ID

	// Don't talk to oneself
	if author == myselfID {
		return
	}

	if msg.GuildID != "" {
		return // DMs only!
	}

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       prettyembedcolor,
		Description: "",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  ":rotating_light: :wheelchair: I have received a DM :wheelchair: :rotating_light:",
				Value: msg.Content,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "from @" + msg.Author.Username + "#" + msg.Author.Discriminator,
		},
	}
	_, err := session.ChannelMessageSendEmbed(FORWARD_TO, embed)
	if err != nil {
		log.Println(err)
	}

	go func() {
		time.Sleep(3 * time.Second)
		SendDM(author, "hey do you need some help?")
	}()
}
