package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const rulesChannel = "667494326372139008"
const rulesMessage = "667497572264312832"

var rules = []string{
	"Moderators have the final say. Do not argue with them.",
	"Use the correct channels (ask questions in <#" + help + ">, report bugs on github, etc)",
	"Channel specific rules or topics can be found in the channel description",
	"No trolling, unnecessary tagging, spamming, NSFW content, bullying, or blatant rudeness",
	"No advertising",
}

var extraRules = []*discordgo.MessageEmbedField{
	{
		Name: "Volunteers",
		Value: "All staff, including Support, Moderators, and Developers are volunteers. " +
			"They are under _no obligation_ to help you, but are likely to if you are polite.",
	},
	{
		Name:  "Why can't I speak‽",
		Value: "You need to verify yourself! Click [here](https://modulobot.xyz/verify/208753003996512258) and follow the prompts. You also need to wait 10 minutes.",
	},
	{
		Name: "Terms",
		Value: "By using this discord you agree to our bots storing information about you, such as your discord id. " +
			"If you wish for this information to be removed from our servers you can run `i!optout` to delete any records we have about you and remove you from the server.",
	},
}

func rulesHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	reply := discordgo.MessageEmbed{
		Color: prettyembedcolor,
	}

	switch len(args) {
	case 1:
		reply.Title = "Rules"
		reply.Description = buildRules()
	case 2:
		index, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		index-- // Rule numbers are one higher than index
		if index >= len(rules) {
			return errors.New("There are only " + strconv.Itoa(len(rules)) + " rules, " + args[1] + " is too high.")
		}
		if index < 0 {
			return errors.New("Rules are counted from 1, " + args[1] + " is too low")
		}
		reply.Title = "Rule " + strconv.Itoa(index+1)
		reply.Description = rules[index]
		reply.Footer = &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Run %s%s for more", prefix, args[0]),
		}
	default:
		return errors.New("incorrect number of arguments")
	}

	_, err := discord.ChannelMessageSendEmbed(msg.ChannelID, &reply)
	return err
}

func updateRules() {
	_, err := discord.ChannelMessageEditEmbed(rulesChannel, rulesMessage, &discordgo.MessageEmbed{
		Color: prettyembedcolor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/attachments/224684271913140224/571442198718185492/unknown.png",
		},
		Fields: append(extraRules, &discordgo.MessageEmbedField{
			Name:  "Rules",
			Value: buildRules(),
		}),
	})
	if err != nil {
		log.Println("Unable to edit rules message with id " + rulesMessage)
	}
}

func buildRules() string {
	var r strings.Builder
	for index, rule := range rules {
		r.WriteString(strconv.Itoa(index + 1))
		r.WriteString(". ")
		r.WriteString(rule)
		r.WriteString("\n")
	}

	return r.String()
}
