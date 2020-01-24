package option_b

import (
	"fmt"
	"strings"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var (
	// Logger is the zap logger
	Logger *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	Logger, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

// Debug prints a log with level debug
func Debug(msg string, tags ...string) {
	fields := parseFields(tags...)
	Logger.Debug(msg, fields...)
	Logger.Sync()
}

// Info prints a log with level Info
func Info(msg string, tags ...string) {
	fields := parseFields(tags...)
	Logger.Info(msg, fields...)
	Logger.Sync()
}

// Error prints a log with level Error
func Error(msg string, err error, tags ...string) {
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	fields := parseFields(tags...)
	Logger.Error(msg, fields...)
	Logger.Sync()
}

// Field translate key value pair into zap field
func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func parseFields(tags ...string) []zap.Field {
	results := make([]zap.Field, len(tags))
	for i, tag := range tags {
		els := strings.Split(tag, ":")
		result := Field(els[0], els[1])
		results[i] = result
	}
	return results
}
