package main

import (
	"github.com/rombintu/checker-sprints/internal/config"
	"github.com/rombintu/checker-sprints/internal/server"
	"github.com/rombintu/checker-sprints/lib/logger"
)

func main() {
	logger.Init(true)
	conf := config.NewServerConfig()
	// s := storage.NewStorage(conf.Driver, conf.Storage.ConnectionString, conf.Storage.Database)
	// defer s.Close()
	// logger.Log.Debug("Storage", slog.Any("open", s.Open()))
	api := server.NewServer(conf)
	api.Configure()
	api.Start()
}
