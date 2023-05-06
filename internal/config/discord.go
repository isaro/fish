package config

import (
	"github.com/bwmarrin/discordgo"
	"github.com/isaro/fish/internal/controller"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "fish",
			Description: "Get a random fish",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"fish": controller.Fish,
	}
)

type DiscordConfig struct {
	anglerToken    string
	DiscordSession *discordgo.Session
}

func initDiscord(discordConfig *DiscordConfig) {
	var err error

	anglerToken := os.Getenv("ANGLER_TOKEN")
	switch "" {
	case anglerToken:
		logrus.Fatal("ANGLER_TOKEN is not set")
	}

	discordConfig.DiscordSession, err = discordgo.New("Bot " + anglerToken)
	if err != nil {
		logrus.Fatal("Error creating Discord session: ", err)
	}

	discordConfig.DiscordSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = discordConfig.DiscordSession.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	//Register discord command handlers
	discordConfig.DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	//Register discord commands
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := discordConfig.DiscordSession.ApplicationCommandCreate(discordConfig.DiscordSession.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

}
