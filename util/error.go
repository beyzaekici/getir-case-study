package util

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func Init() {
	log.Logger = logger
}

func Error(msg error) {
	log.Logger.Error().Err(msg)
}

func Info(msg string) {
	log.Logger.Info().Msg(msg)
}
