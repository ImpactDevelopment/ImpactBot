package main

import (
	"github.com/bwmarrin/discordgo"
)

func giveaway(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	return discord.GuildMemberRoleAdd(IMPACT_SERVER, caller.User.ID, "698619050833477633")
}
