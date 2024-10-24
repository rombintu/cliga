package logger

import (
	"log/slog"
	"os"
	"time"
)

// var Log *slog.Logger

func Init(debug bool) {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: level,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// check that we are handling the time key
				if a.Key != slog.TimeKey {
					return a
				}

				t := a.Value.Time()
				a.Value = slog.StringValue(t.Format(time.RFC822))

				return a
			},
		}))
	slog.SetDefault(logger)
	// if debug {
	// 	slog.SetLogLoggerLevel(slog.LevelDebug)
	// }
}
