package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func onUserJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	if m.GuildID != IMPACT_SERVER {
		return
	}
	if m.User == nil {
		return
	}
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       prettyembedcolor,
		Title:       "Welcome to the Impact Discord!",
		Description: "In order to prevent spam and give you a chance to read the FAQs and rules, you will not be able to talk for ten minutes.\nIn the meantime, check the useful links below. Please do not DM a staff member while waiting. Try to resolve the problem using the FAQ, or the help channel when you can speak.",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Setup/Install FAQ",
				Value:  "[Click here!](https://github.com/ImpactDevelopment/ImpactClient/wiki/Setup-FAQ)",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Usage FAQ",
				Value:  "[Click here!](https://github.com/ImpactDevelopment/ImpactClient/wiki/Usage-FAQ)",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Rules",
				Value:  "[Click here!](https://discordapp.com/channels/208753003996512258/224684271913140224/306183650268020748)",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Github Links",
				Value:  "[Impact](https://github.com/ImpactDevelopment/ImpactClient), [Installer](https://github.com/ImpactDevelopment/Installer/), [Baritone](https://github.com/cabaletta/baritone)",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Tutorial videos for downloading and installing the client",
				Value:  "[Windows](https://www.youtube.com/watch?v=QP6CN-1JYYE)\n[Mac OSX](https://www.youtube.com/watch?v=BBO0v4eq95k)\n[Linux](https://www.youtube.com/watch?v=XPLvooJeQEI)\n",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "♿ Impact Client ♿",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
	ch, err := discord.UserChannelCreate(m.User.ID)
	if err != nil {
		log.Println(err)
		log.Println("Can't begin to send welcome message")
		return
	}
	_, err = discord.ChannelMessageSendEmbed(ch.ID, embed)
	if err != nil {
		log.Println(err)
		log.Println("Can't send welcome message")
		return
	}
}
