package config

import (
	"github.com/sirupsen/logrus"
)

var (
	LogLevel = logrus.InfoLevel
)

func initLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(LogLevel)
}
