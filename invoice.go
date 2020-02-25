package main

import (
	"github.com/bwmarrin/discordgo"
)

func onVoiceStateUpdate(session *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	if m.GuildID != IMPACT_SERVER {
		return
	}
	if m.ChannelID == "" || m.Deaf || m.SelfDeaf {
		session.GuildMemberRoleRemove(IMPACT_SERVER, m.UserID, InVoice.ID)
	} else {
		session.GuildMemberRoleAdd(IMPACT_SERVER, m.UserID, InVoice.ID)
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
