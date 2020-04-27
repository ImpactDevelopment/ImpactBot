package main

import (
	"github.com/bwmarrin/discordgo"
)

func giveaway(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleAdd(impactServer, caller.User.ID, "698619050833477633")
	if err != nil {
		return err
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, ":white_check_mark:")
}
