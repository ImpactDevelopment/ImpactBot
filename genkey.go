package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func genkey(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	if !IsUserAtLeast(caller, Developer) {
		return fmt.Errorf("ahaha epic trolololol moment")
	}

	// Strip command name from args and set default role
	args = args[1:]
	if len(args) < 1 {
		args = []string{"premium"}
	}

	// Serialise args to a role array
	// TODO ignore prefixing args?
	var roles strings.Builder
	for _, role := range args {
		if roles.Len() > 0 {
			roles.WriteString("&")
		}
		roles.WriteString("role=" + url.QueryEscape(role))
	}

	// Default to premium
	if roles.Len() == 0 {
		return errors.New("No valid roles in list")
	}

	// https://api.impactclient.net
	code, err := get("http://api.localhost:3000/v1/integration/impactbot/genkey?auth=" + os.Getenv("IMPACTBOT_AUTH_SECRET") + "&" + roles.String())
	if err != nil {
		return err
	}
	return resp(msg.ChannelID, "Token is `"+code+"`\n[Click here to register an Impact Account](https://impactclient.net/register.html?token="+code+")")
}

func get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(data))
	}
	return string(data), nil
}
