package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func chess(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleAdd(impactServer, caller.User.ID, Chess.ID)
	if err != nil {
		return err
	}
	discord.MessageReactionAdd(msg.ChannelID, msg.ID, check)

	return resp(msg.ChannelID, fmt.Sprintf("User has been given chess role!", caller.User.Username, caller.User.Discriminator))
}

func unchess(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleRemove(impactServer, caller.User.ID, Chess.ID)
	if err != nil {
		return err
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, check)
}
