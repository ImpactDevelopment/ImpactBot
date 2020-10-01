package main

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
)

func handleFunny(_ *discordgo.Member, msg *discordgo.Message, _ []string) error {
	resp, err := http.Get("https://ifunny.co/feeds/shuffle")
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return err
	}
	val, exists := doc.Find(".media__image").First().Attr("data-src")
	if !exists {
		return errors.New("we ran out of runny =(")
	}
	reply := &discordgo.MessageEmbed{
		Title: ":rofl:",
		Image: &discordgo.MessageEmbedImage{
			URL: val,
		},
		Color: prettyembedcolor,
	}
	_, err = discord.ChannelMessageSendEmbed(msg.ChannelID, reply)
	return err
}
