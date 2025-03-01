package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(msg string, fields ...any)
	Error(err error, msg string, fields ...any)
	Debug(msg string, fields ...any)
}

type ZeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger() Logger {
	zerolog.TimeFieldFormat = time.DateTime

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &ZeroLogger{
		logger: logger,
	}
}

func NewFileLogger(filePath string) (Logger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}

	zerolog.TimeFieldFormat = time.DateTime

	logger := zerolog.New(file).With().Timestamp().Logger()
	return &ZeroLogger{
		logger: logger,
	}, nil
}

func (l ZeroLogger) Info(msg string, fields ...any) {
	l.logger.Info().Fields(fields).Msg(msg)
}

func (l ZeroLogger) Error(err error, msg string, fields ...any) {
	l.logger.Error().Fields(fields).Err(err).Msg(msg)
}

func (l ZeroLogger) Debug(msg string, fields ...any) {
	l.logger.Debug().Fields(fields).Msg(msg)
}
