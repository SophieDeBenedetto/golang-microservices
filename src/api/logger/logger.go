package logger

import (
	"os"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/config"
	"github.com/sirupsen/logrus"
)

var (
	// Logger contains the logger
	Logger *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Logger = &logrus.Logger{
		Out:   os.Stdout,
		Level: level,
	}

	if config.IsProduction() {
		Logger.Formatter = &logrus.JSONFormatter{}
	} else {
		Logger.Formatter = &logrus.TextFormatter{}
	}
}
