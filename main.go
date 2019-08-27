package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var discord *discordgo.Session

var myselfID string

const (
	IMPACT_SERVER = "208753003996512258"

	SUPPORT_ROLE = "245682967546953738"

	prettyembedcolor = 3447003
)

func init() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		panic("Must set environment variable DISCORD_BOT_TOKEN")
	}
	log.Println("Establishing discord connection")
	var err error
	discord, err = discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	user, err := discord.User("@me")
	if err != nil {
		panic(err)
	}

	myselfID = user.ID
	log.Println("I am", myselfID)

	discord.AddHandler(onUserJoin)
	discord.AddHandler(onMessageSent)
	discord.AddHandler(onMessageReactedTo)
	discord.AddHandler(onReady)
}

func main() {
	err := discord.Open()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to discord")
	forever := make(chan int)
	<- forever
}
