package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func onUserJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	log.Println("On user join")
	if m.GuildID != impactServer {
		return
	}
	if m.User == nil {
		return
	}
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  prettyembedcolor,
		Title:  "Welcome to the Impact Discord!",
		Fields: append([]*discordgo.MessageEmbedField{
			{
				Name:   "Setup/Install FAQ",
				Value:  "[Click here!](https://github.com/ImpactDevelopment/ImpactClient/wiki/Setup-FAQ)",
				Inline: true,
			},
			{
				Name:   "Usage FAQ",
				Value:  "[Click here!](https://github.com/ImpactDevelopment/ImpactClient/wiki/Usage-FAQ)",
				Inline: true,
			},
			{
				Name:   "Rules",
				Value:  "[Click here!](https://discordapp.com/channels/208753003996512258/667494326372139008/667497572264312832)",
				Inline: true,
			},
			{
				Name:   "Github Links",
				Value:  "[Impact](https://github.com/ImpactDevelopment/ImpactClient), [Installer](https://github.com/ImpactDevelopment/Installer/), [Baritone](https://github.com/cabaletta/baritone)",
				Inline: true,
			},
			{
				Name:  "Tutorial videos for downloading and installing the client",
				Value: "[Windows](https://www.youtube.com/watch?v=QP6CN-1JYYE)\n[Mac OSX](https://www.youtube.com/watch?v=BBO0v4eq95k)\n[Linux](https://www.youtube.com/watch?v=XPLvooJeQEI)\n",
			},
		}, extraRules...),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "♿ Impact Client ♿",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/attachments/224684271913140224/571442198718185492/unknown.png",
		},
	}
	if hasRole(m.Member, Verified) {
		embed.Description = "You have been verified automatically! Check the useful links below or <#" + help + ">"
	} else {
		embed.Description = "**In order to prevent spam you will not be able to talk until you complete a captcha by clicking [this link](https://impactclient.net/discord.html?discord=" + m.User.ID + ")** to prove you're not a bot!\n\nCheck the useful links below. Please do not DM a staff member while waiting. Try to resolve the problem using the FAQ, or the help channel when you can speak."
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
