package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/rombintu/checker-sprints/internal/storage"
)

type StorageConfig struct {
	Database         string
	ConnectionString string
}

type ServerConfig struct {
	Driver  string
	Storage StorageConfig
	Listen  string
}

func tryGetEnv(key string, or string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return or
	}
	return value
}

func NewServerConfig() ServerConfig {
	if err := godotenv.Load(); err != nil {
		slog.Error("env file not loaded", slog.Any("error", err))
	}

	return ServerConfig{
		Driver: tryGetEnv("DRIVER", storage.MemDriverName),
		Storage: StorageConfig{
			ConnectionString: tryGetEnv("CONNECTION", ""),
			Database:         tryGetEnv("DATABASE", "main"),
		},
		Listen: tryGetEnv("ADDRESS", "localhost:8080"),
	}
}
