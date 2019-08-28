package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"syscall"
)

var discord *discordgo.Session

var myselfID string

const (
	IMPACT_SERVER = "208753003996512258"

	SUPPORT_ROLE = "245682967546953738"

	prettyembedcolor = 3447003
)

func init() {
	var err error

	// You can set environment variables in the git-ignored .env file for convenience while running locally
	err = gotenv.Load()
	if err == nil {
		println("Loaded .env file")
	} else if e, ok := err.(*os.PathError); ok && e.Err == syscall.ERROR_FILE_NOT_FOUND {
		println("No .env file found")
		err = nil // Mutating state is bad mkay
	} else {
		panic(err)
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		panic("Must set environment variable DISCORD_BOT_TOKEN")
	}
	log.Println("Establishing discord connection")
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
	<-forever
}
