package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	logLevel, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	log.SetLevel(logrus.Level(logLevel))
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	return log
}
