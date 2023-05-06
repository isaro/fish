package config

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	anglerToken    string
	discordSession *discordgo.Session
)

func initDiscord() {
	anglerToken = os.Getenv("ANGLER_TOKEN")
	discordSession, err := discordgo.New("Bot " + anglerToken)
	if err != nil {
		logrus.Fatal(err)
	}

}
