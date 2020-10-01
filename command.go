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
	Aliases     []string
	Description string
	Usage       []string
	RoleNeeded  *Role
	Handler     func(caller *discordgo.Member, message *discordgo.Message, args []string) error
}

// List of commands
var Commands = []Command{
	{
		Name:        "optout",
		Description: "Opt out of our terms and leave the server permanently",
		Usage:       []string{"i am sure"},
		Handler:     optOutHandler,
	},
	{
		Name:        "tempmute",
		Description: "Mute someone temporarily, optionally from a specific channel",
		Usage:       []string{"@user reason", "@user #channel reason", "#channel @user reason"},
		RoleNeeded:  &Support,
		Handler:     muteHandler,
	},
	{
		Name:        "mute",
		Description: "Mute someone permanently, optionally from a specific channel",
		Usage:       []string{"@user reason", "@user #channel reason", "#channel @user reason"},
		RoleNeeded:  &Moderator,
		Handler:     muteHandler,
	},
	{
		Name:        "unmute",
		Description: "Unmute someone, either server-wide, from a specific channel, or remove all mutes",
		Usage:       []string{"@user", "@user #channel", "@user all"},
		RoleNeeded:  &Moderator,
		Handler:     unmuteHandler,
	},
	{
		Name:        "kick",
		Description: "Kick someone from the server",
		Usage:       []string{"@user reason"},
		RoleNeeded:  &Moderator,
		Handler:     rektHandler,
	},
	{
		Name:        "ban",
		Description: "Ban someone from the server",
		Usage:       []string{"@user reason"},
		RoleNeeded:  &Moderator,
		Handler:     rektHandler,
	},
	{
		Name:        "rules",
		Aliases:     []string{"rule"},
		Description: "Display the rules or a specific rule. Optionally @mention a user to tag them in the response.",
		Usage: []string{
			"",
			"<number>",
			"<search term>",
		},
		Handler: rulesHandler,
	},
	/*
		{
			Name:        "want",
			Description: "want a nick",
			Usage: []string{
				"<number>",
			},
			RoleNeeded: &Support,
			Handler:    wantHandler,
		},
	*/
	{
		Name:        "cringe",
		Description: "Generates a random cringe image",
		Usage:       []string{""},
		Handler:     handleCringe,
	},
	{
		Name:        "addcringe",
		Description: "Adds a cringe photo to the collection",
		Usage:       []string{"", "url"},
		RoleNeeded:  &Support,
		Handler:     handleAddCringe,
	},
	{
		Name:        "delcringe",
		Description: "Removes a cringe photo from the collection",
		Usage:       []string{"", "url"},
		RoleNeeded:  &Moderator,
		Handler:     handleDelCringe,
	},
	{
		Name:        "genkey",
		Description: "Generates an Impact premium key",
		Usage:       []string{"", "role [...roles]"},
		RoleNeeded:  &SeniorMod,
		Handler:     genkey,
	},
	{
		Name:        "giveaway",
		Description: "Gives you the giveaway role",
		Usage:       []string{""},
		Handler:     giveaway,
	},
	{
		Name:        "ungiveaway",
		Description: "Removes the giveaway role",
		Usage:       []string{""},
		Handler:     ungiveaway,
	},
	{
		Name:        "stupid",
		Description: "makes you so stupid impcat bot will ignore you",
		Usage:       []string{""},
		Handler:     stupid,
	},
	{
		Name:        "unstupid",
		Description: "no more stupid",
		Usage:       []string{""},
		Handler:     unstupid,
	},
	{
		Name:        "funny",
		Aliases:     []string{"unfunny"},
		Description: "didnt laugh",
		Usage:       []string{""},
		Handler:     handleFunny,
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
		for _, alias := range it.Aliases {
			if command == alias {
				return &it
			}
		}
	}
	return nil
}

var helpCommand = Command{
	Name:        "help",
	Description: "Display this help message. Specify `all` to include commands you don't have permission to run. Specify a command's name or alias to see help for only that specific command.",
	Aliases:     []string{"?"},
	Usage: []string{
		"",
		"all",
		"<command>",
	},
	Handler: func(caller *discordgo.Member, message *discordgo.Message, args []string) error {
		embed := discordgo.MessageEmbed{
			Color:  prettyembedcolor,
			Fields: []*discordgo.MessageEmbedField{},
		}
		// all is true if the user asked for commands they don't have permission for
		all := len(args) == 2 && strings.ToLower(args[1]) == "all"
		if len(args) < 2 || all {
			// All commands
			embed.Title = "ImpactBot help"
			embed.Description = "Available commands:"
			if all {
				embed.Description = "All commands:"
			}
			for _, command := range Commands {
				// Include this command if the user has permission to run it or they asked for "all" commands
				if all || command.RoleNeeded == nil || IsUserAtLeast(caller, *command.RoleNeeded) {
					embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
						Name:  command.Name,
						Value: command.helpText(),
					})
				}
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
	// Print aliases first
	if len(c.Aliases) > 0 {
		desc.WriteString("\n_Alias")
		if len(c.Aliases) > 1 {
			// plural meme
			desc.WriteString("es")
		}
		desc.WriteString(": ")
		for _, alias := range c.Aliases {
			desc.WriteString(fmt.Sprintf("**%s** ", alias))
		}
		desc.WriteString("_\n")
	}
	// Then usages
	if len(c.Usage) > 0 {
		desc.WriteString("```\n")
		for _, usage := range c.Usage {
			desc.WriteString(fmt.Sprintf("%s%s %s\n", prefix, c.Name, usage))
		}
		desc.WriteString("```")
	}
	// Then description
	desc.WriteString(c.Description)
	if !strings.HasSuffix(c.Description, ".") {
		desc.WriteString(".")
	}
	// Then, finally, required permissions
	if c.RoleNeeded != nil {
		desc.WriteString(fmt.Sprintf("\nRequires `%s` or higher", c.RoleNeeded.Name))
	}

	return desc.String()
}
