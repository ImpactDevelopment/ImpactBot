package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func genkey(caller *discordgo.Member, msg *discordgo.Message, _ []string) error {
	if !IsUserAtLeast(caller, Developer) {
		return fmt.Errorf("ahaha epic trolololol moment")
	}
	return resp(msg.ChannelID, get("https://api.impactclient.net/v1/integration/impactbot/genkey?auth="+os.Getenv("IMPACTBOT_AUTH_SECRET")))
}

func get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
