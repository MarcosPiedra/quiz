package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Level string

type LogConfig struct {
	Environment string
	LogLevel    Level
}

const (
	TRACE Level = "TRACE"
	DEBUG Level = "DEBUG"
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
	PANIC Level = "PANIC"
)

func NewLogger(cfg LogConfig) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = "00:00:00.000PM"
	})).
		Level(logLevelToZero(cfg.LogLevel)).
		With().
		Timestamp().
		Str("env", cfg.Environment).
		Logger()
}

func logLevelToZero(level Level) zerolog.Level {
	switch level {
	case PANIC:
		return zerolog.PanicLevel
	case ERROR:
		return zerolog.ErrorLevel
	case WARN:
		return zerolog.WarnLevel
	case INFO:
		return zerolog.InfoLevel
	case DEBUG:
		return zerolog.DebugLevel
	case TRACE:
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}
