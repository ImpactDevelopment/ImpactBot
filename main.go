package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
)

var discord *discordgo.Session

var myselfID string

const (
	IMPACT_SERVER    = "208753003996512258"
	BRADY            = "205718273696858113"
	DONATOR          = "210114021641289728"
	prettyembedcolor = 3447003
)

var STAFF = map[string]string{
	"headDeveloper": "209817890713632768",
	"developer":     "221655083748687873",
	"seniorModerator": "663065117738663938", // wow so new, it starts with a 6
	"moderator":     "210377982731223040",
	"support":       "245682967546953738",
}

func init() {
	var err error

	// You can set environment variables in the git-ignored .env file for convenience while running locally
	err = gotenv.Load()
	if err == nil {
		println("Loaded .env file")
	} else if os.IsNotExist(err) {
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
	discord.AddHandler(onMessageSent2)
	discord.AddHandler(onMessageSent3)
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
