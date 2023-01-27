package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SystemLog(level zerolog.Level, system string) *zerolog.Event {
	return log.WithLevel(level).Str("system", system)
}
