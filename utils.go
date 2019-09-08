package main

import (
	"github.com/bwmarrin/discordgo"
)

func isStaff(user string) bool {
	for _, role := range STAFF {
		if hasRole(user, role) {
			return true
		}
	}
	return false
}

// True if user has ANY role passed in
func hasRole(user string, role ...string) bool {
	member, err := discord.GuildMember(IMPACT_SERVER, user)
	if err != nil || member == nil {
		return false
	}
	for _, r := range role {
		if includes(member.Roles, r) {
			return true
		}
	}
	return false
}

// True if user has ALL roles passed in
func hasRoles(user string, role ...string) bool {
	member, err := discord.GuildMember(IMPACT_SERVER, user)
	if err != nil || member == nil {
		return false
	}
	for _, r := range role {
		if !includes(member.Roles, r) {
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

func SendDM(user_id string, message string) error {
	ch, err := discord.UserChannelCreate(user_id) // only creates it if it doesn"t already exist
	if err != nil {
		return err
	}
	_, err = discord.ChannelMessageSend(ch.ID, message)
	return err
}
