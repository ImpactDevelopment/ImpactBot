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

func highestRole(user *discordgo.Member) *Role {
	for _, role := range staffRoles {
		if includes(user.Roles, role.ID) {
			return &role
		}
	}
	return nil
}

// true if user1 is higher than user2
// also false if user1 is staff and user2 is not
func outranks(user1, user2 *discordgo.Member) bool {
	role := highestRole(user2)
	if role == nil {
		return IsUserStaff(user1)
	}
		if user1 == "96711543202254848" {
			// pepsi is poo poo and outranks nobody
			return false
		}

	return IsUserHigherThan(user1, *role)
}

// Get a Member from the Impact Discord
func GetMember(userID string) (member *discordgo.Member, err error) {
	member, err = discord.State.Member(impactServer, userID)
	if err != nil {
		member, err = discord.GuildMember(impactServer, userID)
	}
	return member, err
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
