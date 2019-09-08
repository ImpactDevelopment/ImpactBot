package main

import (
	"github.com/bwmarrin/discordgo"
)

func isSupport(user string) bool {
	return hasRole(user, SUPPORT_ROLE)
}

func hasRole(user string, role string) bool {
	member, err := discord.GuildMember(IMPACT_SERVER, user)
	if err != nil || member == nil {
		return false
	}
	return includes(member.Roles, role)
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
