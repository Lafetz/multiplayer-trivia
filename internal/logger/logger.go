package logger

import (
	"log/slog"
	"os"
)

var LogLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func NewLogger(level string, env string) *slog.Logger {

	logLevel := LogLevels[level]

	var logHandler slog.Handler

	switch env {
	case "development":
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	default:
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		})

	}
	logger := slog.New(logHandler)
	return logger
}
