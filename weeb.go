package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func optOutHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) (err error) {
	if len(args) < 2 || strings.ToLower(strings.Join(args[1:], " ")) != "yes" {
		return fmt.Errorf("You will be **permanently banned from this server** by informing us that you are a weeb. Are you sure? `%s%s %s`", prefix, args[0], "yes")
	}
		_ = SendDM(caller.User.ID, "no anime allowed on this server.")
		err = discord.GuildBanCreateWithReason(impactServer, caller.User.ID, "no weebs allowed.", 0)
		if err != nil {
			return fmt.Errorf("We were unable to ban you. Please contact a moderator.\nError: %s", err.Error())
		}
	return resp(msg.ChannelID, fmt.Sprintf("User @%s#%s has informed us that they are a weeb and has now been banned off of the server.", caller.User.Username, caller.User.Discriminator))
}
