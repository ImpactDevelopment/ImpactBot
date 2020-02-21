package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onReady2(discord *discordgo.Session, ready *discordgo.Ready) {
	go func() {
		prev := ""
		total := 0
		for {
			st, err := discord.GuildMembers(IMPACT_SERVER, prev, 1000)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("Fetched", len(st), "more members, total so far is", total)
			if len(st) == 0 {
				log.Println("No more members!")
				break
			}
			prev = st[len(st)-1].User.ID
			for _, member := range st {
				total++
				handlePoorBaby(member)
			}
		}
		log.Println("Processed", total, "members")
	}()
}

func handlePoorBaby(member *discordgo.Member) {
	if hasRole(member, Donator) && !hasRole(member, Verified) {
		log.Println("Member", member.User.ID, "had donator but not verified")
		err := discord.GuildMemberRoleAdd(IMPACT_SERVER, member.User.ID, Verified.ID)
		if err != nil {
			log.Println(err)
		}
	}
}
