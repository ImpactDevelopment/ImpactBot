package main
// Naive, bad and untested, just a modified weeb
// I literally read 2 lines of documentation for this, don't expect good shit
import (
	"github.com/bwmarrin/discordgo"
)

func unweeb(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleDelete(impactServer, caller.User.ID, "612744883467190275")
	if err != nil {
		return err
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
