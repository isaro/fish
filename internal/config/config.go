package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	once sync.Once
	cfg  *config
)

type config struct {
	environment   string
	discordConfig DiscordConfig
}

func GetConfig() *config {
	once.Do(func() {
		cfg = setup()
		initLogger()
		initDiscord(&cfg.discordConfig)
	})
	return cfg
}

func setup() (c *config) {
	c = &config{}
	c.discordConfig = DiscordConfig{}
	c.environment = os.Getenv("ENVIRONMENT")

	logrus.Debug("Environment: ", c.Environment)

	return c
}

func (c config) Environment() string {
	return c.environment
}

func (c config) DiscordConfig() DiscordConfig {
	return c.discordConfig
}
