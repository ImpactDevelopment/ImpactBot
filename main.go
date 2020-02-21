package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
)

var discord *discordgo.Session

var myselfID string

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
	discord.AddHandler(onMessageSentCommandHandler)
	discord.AddHandler(onUserJoin2)
	discord.AddHandler(onMessageSent3)
	discord.AddHandler(onMessageUpdate)
	discord.AddHandler(onVoiceStateUpdate)
	discord.AddHandler(onUserJoin3)
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
	log.Printf("Impcat bot has started on %d servers:", len(servers))
	for _, guild := range servers {
		log.Println("Server ID", guild.ID, "Name", guild.Name)
	}

	// Replace rules message
	updateRules()
}
