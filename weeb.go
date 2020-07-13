  
package main

import (
	"github.com/bwmarrin/discordgo"
)

func weeb(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleAdd(impactServer, caller.User.ID, "612744883467190275")
	if err != nil {
		return err
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
