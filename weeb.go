  
package main

import (
	"github.com/bwmarrin/discordgo"
	"utils.go"
)

func weeb(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildBanCreateWithReason(impactServer, caller.User.ID, "no weeb for you, poggers.", 0) 
	
	if err != nil {
		return err // if this happens i fucked up
	}
	
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
