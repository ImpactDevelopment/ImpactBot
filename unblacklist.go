package main
// Removes the blacklist role
import (
	"github.com/bwmarrin/discordgo"
)

var blacklistRole = "000000000000000000"

func unblacklist(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := GuildMemberRoleRemove(impactServer, caller.User.ID, blacklistRole)
	if err != nil {
		return err
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
