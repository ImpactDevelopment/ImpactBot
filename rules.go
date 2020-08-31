package main

import (
	"errors"
	"fmt"
	stripmd "github.com/writeas/go-strip-markdown"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const rulesChannel = "667494326372139008"
const rulesMessage = "667497572264312832"

var rules = []string{
	"Follow [Discord's ToS](https://discord.com/terms).",
	"Moderators have the final say. Do not argue with them.",
	"Use the correct channel for the topic you are discussing. _Ask questions in <#" + help + ">, report bugs on [GitHub](https://github.com/ImpactDevelopment/ImpactIssues/issues), etc._",
	"No trolling, unnecessary pinging / @-ing / DMing, spamming, advertising, non-English conversation, bullying, or blatant rudeness.",
	"No NSFW content. _This includes messages, images, gifs, videos, audio, links, usernames, nicknames, statuses, profile pictures, etc._",
	"Don't “ask to ask”. _This means no messages that just say something like “Can I ask a question?”, or “Can someone help me?”, or, worst of all, “hello??”. The answer is yes, in <#" + help + ">)._",
	"Channel specific rules or topics can be found in the channel description. _They may extend upon or overrule these blanket rules._",
}

var extraRules = []*discordgo.MessageEmbedField{
	{
		Name: "Volunteers",
		Value: "All staff, including Support, Moderators, and Developers are volunteers. " +
			"They are under _no obligation_ to help you, but are likely to if you are polite.",
	},
	{
		Name:  "Why can't I speak‽",
		Value: "You need to verify yourself! Click the link at the top of my welcome DM to you, if you have it. Otherwise, click [here](https://impactclient.net/discord.html?member=1) and fill in your info manually.",
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

	// Separate @mentions from args
	var b strings.Builder
	if len(args) > 0 {
		for _, word := range args[1:] {
			if match, err := regexp.MatchString("^<@!?[0-9]+>$", word); err != nil {
				return err
			} else if !match {
				b.WriteString(word + " ")
			}
		}
	}

	// Try to parse an index or a search term
	search := strings.TrimSpace(b.String())
	index, _ := strconv.Atoi(search)

	if search == "" {
		reply.Title = "Rules"
		reply.Description = buildRules()
	} else {
		// If search matches index, then they specified a rule number
		// otherwise they provided a search term
		if search == strconv.Itoa(index) {
			index-- // Rule numbers are one higher than index
		} else {
			var err error
			index, err = findRuleFromStrings(strings.TrimSpace(search))
			if err != nil {
				return err
			}
		}

		if index >= len(rules) {
			return errors.New("There are only " + strconv.Itoa(len(rules)) + " rules, " + strconv.Itoa(index+1) + " is too high.")
		}
		if index < 0 {
			return errors.New("Rules are counted from 1, " + strconv.Itoa(index+1) + " is too low")
		}

		reply.Title = "Rule " + strconv.Itoa(index+1)
		reply.Description = rules[index]
		reply.Footer = &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Run %s%s for more", prefix, args[0]),
		}
	}

	// Mention any users mentioned by the caller
	var mentions string
	for _, user := range msg.Mentions {
		mentions += user.Mention() + " "
	}

	_, err := discord.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
		Content: mentions,
		Embed:   &reply,
	})

	return err
}

func findRuleFromStrings(phrase ...string) (int, error) {
	for _, word := range phrase {
		for i, rule := range rules {
			if strings.Contains(strings.ToLower(stripmd.Strip(rule)), strings.ToLower(stripmd.Strip(word))) {
				return i, nil
			}
		}
	}
	return -1, errors.New("unable to find rule matching \"" + strings.Join(phrase, " ") + "\"")
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
