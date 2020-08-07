package main

import (
	"github.com/bwmarrin/discordgo"
)

func blacklist(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleAdd(impactServer, caller.User.ID, "000000000000000000")
	if err != nil {
		return err
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
