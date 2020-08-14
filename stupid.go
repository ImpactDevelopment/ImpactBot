package main

import (
	"github.com/bwmarrin/discordgo"
)

func stupid(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleAdd(impactServer, caller.User.ID, "743903534160019476")
	if err != nil {
		return err // fuckups get caught here
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
