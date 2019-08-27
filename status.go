package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onReady(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Game: &discordgo.Game{
			Name: "the Impact Discord",
			Type: discordgo.GameTypeWatching,
		},
		AFK:    false,
		Status: "",
	})
	if err != nil {
		log.Println("Error attempting to set my status")
		log.Println(err)
	}
	servers := discord.State.Guilds
	log.Printf("Impcat bot has started on %d servers", len(servers))
}
