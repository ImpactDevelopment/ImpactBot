// Part of ImapctBot
package main
// Removes the giveaway role
import (
	"github.com/bwmarrin/discordgo"
)

var giveawayRole = "698619050833477633"

func ungiveaway(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := GuildMemberRoleRemove(impactServer, caller.User.ID, giveawayRole)
	if err != nil {
		return err
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
