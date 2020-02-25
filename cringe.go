package main

import (
	"fmt"
	"net/url"

	"github.com/bwmarrin/discordgo"
)

// this file is the story of my life lol

func handleCringe(_ *discordgo.Member, msg *discordgo.Message, _ []string) error {
	var rngCringe string
	err := DB.QueryRow("SELECT image FROM cringe ORDER BY RANDOM() LIMIT 1").Scan(&rngCringe)
	if err != nil {
		return err
	}
	reply := discordgo.MessageEmbed{
		Title: ":camera_with_flash:",
		Image: &discordgo.MessageEmbedImage{
			URL: rngCringe,
		},
		Color: prettyembedcolor,
	}
	_, err = discord.ChannelMessageSendEmbed(msg.ChannelID, &reply)
	return err
}

func handleAddCringe(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	if !IsUserAtLeast(caller, Support) {
		return fmt.Errorf("you have to be at least support to call something cringe-worthy lol")
	}

	if len(args) < 2 {
		if len(msg.Attachments) > 0 {
			_, err := DB.Exec("INSERT INTO cringe(image) VALUES($1)", msg.Attachments[0].URL)
			if err == nil {
				return discord.MessageReactionAdd(msg.ChannelID, msg.ID, stevePig)
			}
			return err
		}
		return fmt.Errorf("error : no attachments / links found to add")
	}
	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		return fmt.Errorf("invalid url scheme")
	}
	_, err = DB.Exec("INSERT INTO cringe(image) VALUES($1)", args[1])
	if err == nil {
		return discord.MessageReactionAdd(msg.ChannelID, msg.ID, stevePig)
	}
	return err
}
