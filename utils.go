package main

import (
	"github.com/bwmarrin/discordgo"
)

func isStaff(user *discordgo.Member) bool {
	for _, role := range STAFF {
		if hasRole(user, role) {
			return true
		}
	}
	return false
}

// True if user has ANY role passed in
func hasRole(user *discordgo.Member, role ...string) bool {
	for _, r := range role {
		if includes(user.Roles, r) {
			return true
		}
	}
	return false
}

// True if user has ALL roles passed in
func hasRoles(user *discordgo.Member, role ...string) bool {
	for _, r := range role {
		if !includes(user.Roles, r) {
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
