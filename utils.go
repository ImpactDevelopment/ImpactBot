package main

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

// True if user has ANY role passed in
func hasRole(user *discordgo.Member, role ...Role) bool {
	for _, r := range role {
		if includes(user.Roles, r.ID) {
			return true
		}
	}
	return false
}

// True if user has ALL roles passed in
func hasRoles(user *discordgo.Member, role ...Role) bool {
	for _, r := range role {
		if !includes(user.Roles, r.ID) {
			return false
		}
	}
	return true
}

func mentionsMe(msg *discordgo.Message) bool {
	for _, user := range msg.Mentions {
		if user != nil && user.ID == myselfID {
			return true
		}
	}
	return false
}

func includes(list []string, val string) bool {
	for _, x := range list {
		if x == val {
			return true
		}
	}
	return false
}

func SendDM(userID string, message string) error {
	ch, err := discord.UserChannelCreate(userID) // only creates it if it doesn"t already exist
	if err != nil {
		return err
	}
	_, err = discord.ChannelMessageSend(ch.ID, message)
	return err
}

// Get a Member from the Impact Discord
func GetMember(userID string) (member *discordgo.Member, err error) {
	return discord.GuildMember(IMPACT_SERVER, userID)
}

func findNamedMatches(r *regexp.Regexp, str string) map[string]string {
	matches := r.FindStringSubmatch(str)
	names := r.SubexpNames()
	subs := map[string]string{}

	for i, sub := range matches {
		if names[i] != "" {
			subs[names[i]] = sub
		}
	}

	return subs
}
