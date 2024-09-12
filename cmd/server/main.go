package main

import (
	"github.com/rombintu/checker-sprints/internal/config"
	"github.com/rombintu/checker-sprints/internal/server"
	"github.com/rombintu/checker-sprints/internal/storage"
	"github.com/rombintu/checker-sprints/lib/logger"
)

func main() {
	logger.Init(true)
	conf := config.NewServerConfig()
	store := storage.NewStorage(conf.Driver, conf.Storage.ConnectionString, conf.Storage.Database)
	api := server.NewServer(conf, store)
	api.Configure()
	api.Start()
}
