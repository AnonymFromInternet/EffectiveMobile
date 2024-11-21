package logger

import (
	"log/slog"
	"os"
)

// Levels: error, debug, info

const (
	LOCAL = "LOCAL"
	DEBUG = "DEBUG"
	PROD  = "PROD"
)

func MustCreate(mode string) *slog.Logger {
	var logger *slog.Logger

	switch mode {
	case LOCAL:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case DEBUG:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case PROD:
		file, e := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if e != nil {
			logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
			logger.Error("cannot open log file")
		} else {
			logger = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo}))
		}
	}

	return logger
}
