package logger

import (
	"log/slog"
	"os"
	"wb-data-service-golang/wb-data-worker/internal/domain"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type _Logger struct {
	Slogger slog.Logger
}

func New(slogger *slog.Logger) domain.Logger {
	return _Logger{
		Slogger: *slogger,
	}
}

func (logger _Logger) Debug(msg string, args domain.LoggerArgs) {
	logger.Slogger.Debug(msg, args)
}

func (logger _Logger) Error(err error, args domain.LoggerArgs) {
	logger.Slogger.Error(err.Error(), args)
}

func (logger _Logger) Info(msg string, args domain.LoggerArgs) {
	logger.Slogger.Info(msg, args)
}

func (logger _Logger) Warn(msg string, args domain.LoggerArgs) {
	logger.Slogger.Warn(msg, args)
}

func NewLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		panic("Incorrect environment for logger: " + env)
	}

	return log
}
