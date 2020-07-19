package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix        string
	prefixPattern *regexp.Regexp
)

type Command struct {
	Name        string
	Description string
	Usage       []string
	RoleNeeded  *Role
	Handler     func(caller *discordgo.Member, message *discordgo.Message, args []string) error
}

// List of commands
var Commands = []Command{
	{
		Name:        "optout",
		Description: "opt-out of out ToS and leave the server permanently",
		Usage:       []string{"i am sure"},
		Handler:     optOutHandler,
	},
	{
		Name:        "tempmute",
		Description: "mute someone temporarily, optionally from a specific channel",
		Usage:       []string{"@user reason", "@user #channel reason", "#channel @user reason"},
		RoleNeeded:  &Support,
		Handler:     muteHandler,
	},
	{
		Name:        "mute",
		Description: "mute someone permanently, optionally from a specific channel",
		Usage:       []string{"@user reason", "@user #channel reason", "#channel @user reason"},
		RoleNeeded:  &Moderator,
		Handler:     muteHandler,
	},
	{
		Name:        "unmute",
		Description: "unmute someone, either server-wide, from a specific channel, or remove all mutes",
		Usage:       []string{"@user", "@user #channel", "@user all"},
		RoleNeeded:  &Moderator,
		Handler:     unmuteHandler,
	},
	{
		Name:        "kick",
		Description: "kick someone from the server",
		Usage:       []string{"@user reason"},
		RoleNeeded:  &Moderator,
		Handler:     rektHandler,
	},
	{
		Name:        "ban",
		Description: "ban someone from the server",
		Usage:       []string{"@user reason"},
		RoleNeeded:  &Moderator,
		Handler:     rektHandler,
	},
	{
		Name:        "rules",
		Description: "display the rules",
		Usage: []string{
			"",
			"<number>",
		},
		Handler: rulesHandler,
	},
	{
		Name:        "want",
		Description: "want a nick",
		Usage: []string{
			"<number>",
		},
		RoleNeeded:  &Support,
		Handler: wantHandler,
	},
	{
		Name:        "cringe",
		Description: "generates a random cringe image",
		Usage:       []string{""},
		Handler:     handleCringe,
	},
	{
		Name:        "addcringe",
		Description: "adds a cringe photo to the collection",
		Usage:       []string{"", "url"},
		RoleNeeded:  &Support,
		Handler:     handleAddCringe,
	},
	{
		Name:        "delcringe",
		Description: "removes a cringe photo from the collection",
		Usage:       []string{"", "url"},
		RoleNeeded:  &Moderator,
		Handler:     handleDelCringe,
	},
	{
		Name:        "genkey",
		Description: "generates an Impact premium key",
		Usage:       []string{"", "role [...roles]"},
		RoleNeeded:  &SeniorMod,
		Handler:     genkey,
	},
	{
		Name:        "giveaway",
		Description: "gives you the giveaway role",
		Usage:       []string{""},
		Handler:     giveaway,
	},
	{
		Name:        "weeb",
		Description: "gives you the fucking weeb role",
		Usage:       []string{""},
		Handler:     weeb,
	},
}

func init() {
	// Load prefix from the environment
	prefix = os.Getenv("IMPACT_PREFIX")
	if prefix == "" {
		prefix = "i!"
	}
	// Match case-insensitive & ignore whitespace around prefix
	prefixPattern = regexp.MustCompile(`(?i)^\s*` + regexp.QuoteMeta(prefix) + `\s*`)

	// Have to append helpCommand after initializing Commands to avoid an initialization loop
	Commands = append(Commands, helpCommand)
}

func onMessageSentCommandHandler(session *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault || msg.Author.ID == myselfID {
		return // wtf
	}
	if msg.GuildID != impactServer && msg.GuildID != "" {
		return // Only allow guild messages and DMs
	}

	content := msg.Content
	content = strings.Replace(content, ">", "> ", -1)
	content = strings.Replace(content, "<", " <", -1)

	if match := prefixPattern.FindString(content); match != "" { // bot woke
		args := strings.Fields(content[len(match):])
		command := findCommand(strings.ToLower(args[0]))
		if command == nil {
			_ = resp(msg.ChannelID, fmt.Sprintf("Command \"%s\" not found! Try %shelp", args[0], prefix))
			return
		}
		author, err := GetMember(msg.Author.ID)
		if err != nil {
			return
		}
		if command.RoleNeeded != nil && !IsUserAtLeast(author, *command.RoleNeeded) {
			_ = resp(msg.ChannelID, fmt.Sprintf("Command \"%s\" requires at least %s", command.Name, command.RoleNeeded.Name))
			return
		}
		err = command.Handler(author, msg, args)
		if err != nil {
			_ = resp(msg.ChannelID, fmt.Sprintf("Command \"%s\" returned an error: %s", command.Name, err.Error()))
			return
		}

	}
}

func findCommand(command string) *Command {
	for _, it := range Commands {
		if command == it.Name {
			return &it
		}
	}
	return nil
}

var helpCommand = Command{
	Name:        "help",
	Description: "display this help message",
	Usage:       nil,
	RoleNeeded:  nil,
	Handler: func(caller *discordgo.Member, message *discordgo.Message, args []string) error {
		embed := discordgo.MessageEmbed{
			Color:  0,
			Fields: []*discordgo.MessageEmbedField{},
		}
		if len(args) < 2 {
			// All commands
			embed.Title = "ImpactBot help"
			embed.Description = "Available commands:"
			for _, command := range Commands {
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:  command.Name,
					Value: command.helpText(),
				})
			}
		} else {
			// Specified commands
			command := findCommand(args[1])
			if command == nil {
				return errors.New(fmt.Sprintf("Command \"%s\" not found! Try %shelp", args[1], prefix))
			}
			embed.Title = fmt.Sprintf("Command `%s` usage", command.Name)
			embed.Description = command.helpText()
			embed.Footer = &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("See %shelp for alternative commands", prefix),
			}
		}
		_, err := discord.ChannelMessageSendEmbed(message.ChannelID, &embed)
		if err != nil {
			return err
		}

		return nil
	},
}

func (c Command) helpText() string {
	var desc strings.Builder
	if len(c.Usage) > 0 {
		desc.WriteString("```\n")
		for _, usage := range c.Usage {
			desc.WriteString(fmt.Sprintf("%s%s %s\n", prefix, c.Name, usage))
		}
		desc.WriteString("```")
	}
	desc.WriteString(c.Description)
	if !strings.HasSuffix(c.Description, ".") {
		desc.WriteString(".")
	}
	if c.RoleNeeded != nil {
		desc.WriteString(fmt.Sprintf("\nRequires `%s` or higher", c.RoleNeeded.Name))
	}

	return desc.String()
}
