package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	FORWARD_TO   = "617549691730526209"
	NO_HELP_ROLE = "230803433752363020"
)

func canDMBot(userId string) bool {
	member, err := discord.GuildMember(IMPACT_SERVER, userId)
	if err != nil || member == nil {
		return false
	}
	return !includes(member.Roles, NO_HELP_ROLE)
}

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

	if !canDMBot(author) {
		return
	}

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       prettyembedcolor,
		Description: "",
		Fields: []*discordgo.MessageEmbedField{
			{
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
		time.Sleep(1 * time.Second)
		ch, err := discord.UserChannelCreate(author)
		if err != nil {
			log.Println(err)
			return
		}
		err = discord.ChannelTyping(ch.ID)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(2 * time.Second)

		_, err = discord.ChannelMessageSend(ch.ID, fmt.Sprintf("Hey, do you need some help? Checkout <#%s>", help))
		if err != nil {
			log.Println(err)
		}
	}()
}
