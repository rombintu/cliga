package main

import (
	"log/slog"

	"github.com/rombintu/checker-sprints/internal/config"
	"github.com/rombintu/checker-sprints/lib/logger"
)

func main() {
	logger.Init(slog.LevelDebug)
	conf := config.NewCliConfig()
	logger.Log.Debug("Config", slog.Any("data", conf))

	if !conf.Debug {
		logger.Init(slog.LevelInfo)
	}

	conf.ParseFromFlags()
	logger.Log.Debug("Config", slog.Any("data", conf))

}
