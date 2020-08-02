package main

import (
	"github.com/bwmarrin/discordgo"
)

func onVoiceStateUpdate(session *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	if m.GuildID != impactServer {
		return
	}
	if m.ChannelID == "" || m.Deaf || m.SelfDeaf {
		_ = session.GuildMemberRoleRemove(impactServer, m.UserID, InVoice.ID)
	} else {
		_ = session.GuildMemberRoleAdd(impactServer, m.UserID, InVoice.ID)
	}
}

func checkDeservesInVoiceRole(userid string) bool {
	for _, guild := range discord.State.Guilds {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == userid && !vs.Deaf && !vs.SelfDeaf {
				return true
			}
		}
	}
	return false
}
