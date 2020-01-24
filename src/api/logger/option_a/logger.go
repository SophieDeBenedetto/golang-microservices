package logger

import (
	"fmt"
	"os"
	"strings"

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
		Logger.Formatter = &logrus.JSONFormatter{}
		// Logger.Formatter = &logrus.TextFormatter{}
	}
}

// Info logs at the info level
func Info(msg string, tags ...string) {
	fmt.Println(tags)
	if Logger.Level < logrus.InfoLevel {
		return
	}
	Logger.WithFields(parseFields(tags...)).Info(msg)
}

// Error logs at the info level
func Error(msg string, err error, tags ...string) {
	if Logger.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Logger.WithFields(parseFields(tags...)).Error(msg)
}

// Debug logs at the info level
func Debug(msg string, tags ...string) {
	if Logger.Level < logrus.DebugLevel {
		return
	}
	Logger.WithFields(parseFields(tags...)).Debug(msg)
}

func parseFields(tags ...string) logrus.Fields {
	results := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		results[els[0]] = els[1]
	}
	return results
}
