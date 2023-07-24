package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

type ZeroLogger struct {
	logger zerolog.Logger
}

var ZeroLog ZeroLogger

// Init
func yrdy(level string, output string) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}

	var err error

	if output == "stdout" && os.Getenv("ENV") != "production" {
		ZeroLog.logger = zerolog.New(zerolog.ConsoleWriter{
			Out: os.Stderr, NoColor: false,
		}).With().Timestamp().Logger()
	} else {
		if output == "stdout" {
			ZeroLog.logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
		} else if output == "stderr" {
			ZeroLog.logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
		} else {
			f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return err
			}
			ZeroLog.logger = zerolog.New(f).With().Timestamp().Logger()
		}
	}
	return err
}

// Debug logs a debug message.
func Debug(msg string) {
	ZeroLog.logger.Debug().Msg(msg)
}

// Info logs an info message.
func Info(msg string) {
	ZeroLog.logger.Info().Msg(msg)
}

// Info logs an info message.
func Infof(format string, v ...interface{}) {
	ZeroLog.logger.Info().Msgf(format, v...)
}

// Warn logs a warning message.
func Warn(msg string) {
	ZeroLog.logger.Warn().Msg(msg)
}

// Error logs an error message.
func Error(msg string, err error) {
	ZeroLog.logger.Err(err).Msg(msg)
}

// Fatal logs a fatal message and exits the program.
func Fatal(msg string, err error) {
	ZeroLog.logger.Fatal().Err(err).Msg(msg)
}

// Panic logs a panic message and panics.
func Panic(msg string, err error) {
	ZeroLog.logger.Panic().Err(err).Msg(msg)
}
