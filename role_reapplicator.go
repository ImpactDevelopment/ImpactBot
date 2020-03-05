package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

// note: this is not for mutes. rekt does that
// this is for roles determined by the central db
func onUserJoin3(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	if m.GuildID != IMPACT_SERVER {
		return
	}
	if shouldGiveDonator(m.User.ID) {
		err := discord.GuildMemberRoleAdd(IMPACT_SERVER, m.User.ID, Donator.ID)
		if err != nil {
			log.Println(err)
		}
	}
}

func shouldGiveDonator(discordID string) bool {
	url := "https://api.impactclient.net/v1/integration/impactbot/checkdonator/" + discordID + "?auth=" + os.Getenv("IMPACTBOT_AUTH_SECRET")
	return isYes(url)
}

func isYes(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	return string(data) == "yes"
}
