  
package main

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

func weeb(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
    if rand.Intn(1) == 0 {
    	err := discord.GuildBanCreateWithReason(impactServer, caller.User.ID, "no weeb for you, poggers.", 0) 	
    }

    else {
    	err := discord.GuildMemberRoleRemove(impactServer, caller.User.ID, "612744883467190275")
    }
    
	if err != nil {
		return err // if this happens i fucked up
	}
	
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
