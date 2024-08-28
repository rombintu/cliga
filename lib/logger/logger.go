package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func Init(level slog.Level) {
	Log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(Log)
}
