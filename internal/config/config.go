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
	environment string
}

func GetConfig() *config {
	once.Do(func() {
		initLogger()
		cfg = setup()
	})
	return cfg
}

func setup() (c *config) {
	c = &config{}
	c.environment = os.Getenv("ENVIRONMENT")

	logrus.Debug("Environment: ", c.Environment)

	return c
}

func (c config) Environment() string {
	return c.environment
}
