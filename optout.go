package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func optOutHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) (err error) {
	if len(args) < 2 || strings.ToLower(strings.Join(args[1:], " ")) != "i am sure" {
		return fmt.Errorf("You will be **permanently removed from this server (ie, banned)** by continuing. Are you sure? `%s%s %s`", prefix, args[0], "i am sure")
	}

	if DB == nil {
		return errors.New("Unable to connect to database")
	}
	// Delete anything referencing them
	_, err = DB.Exec(`DELETE FROM mutes WHERE discord_id = $1`, caller.User.ID)
	if err != nil {
		return fmt.Errorf("We were unable to clear your info. Please contact a moderator.\nError: %s", err.Error())
	}

	// Check if they're muted
	var muted bool
	for _, role := range muteRoles {
		if hasRole(caller, Role{ID: role}) {
			muted = true
			break
		}
	}

	// If muted ban, if not kick
	if muted {
		// Send them a DM before banning
		_ = SendDM(caller.User.ID, "Thank you for opting out. We have deleted any data we had stored about you. You have been banned from the server to prevent bypassing our moderation system.")

		// We have to ban them or they could bypass a mute
		err = discord.GuildBanCreateWithReason(impactServer, caller.User.ID, "opted out of tos", 0)
		if err != nil {
			return fmt.Errorf("We were unable to ban you. Please contact a moderator.\nError: %s", err.Error())
		}
	} else {
		// Send them a DM before kicking
		_ = SendDM(caller.User.ID, "Thank you for opting out. We have deleted any data we had stored about you and kicked you from the server.")

		err = discord.GuildMemberDeleteWithReason(impactServer, caller.User.ID, "opted out of tos")
		if err != nil {
			return fmt.Errorf("We were unable to ban you. Please contact a moderator.\nError: %s", err.Error())
		}
	}

	return resp(msg.ChannelID, fmt.Sprintf("User @%s#%s has opted out and been banned to prevent mute bypassing", caller.User.Username, caller.User.Discriminator))
}
