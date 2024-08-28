package main

import (
	"log/slog"

	"github.com/rombintu/checker-sprints/internal/config"
	"github.com/rombintu/checker-sprints/internal/storage"
	"github.com/rombintu/checker-sprints/lib/logger"
)

func main() {
	logger.Init(slog.LevelDebug)
	conf := config.NewApiConfig()
	logger.Log.Debug("Config", slog.Any("data", conf))

	s := storage.NewStorage(conf.Driver, conf.Storage.ConnectionString, conf.Storage.Database)
	logger.Log.Debug("Storage", slog.Any("open", s.Open()))
	defer s.Close()
}
