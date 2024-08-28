package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/rombintu/checker-sprints/internal/storage"
	"github.com/rombintu/checker-sprints/lib/logger"
)

type StorageConfig struct {
	Database         string
	ConnectionString string
}

type ApiConfig struct {
	Driver  string
	Storage StorageConfig
}

func tryGetEnv(key string, or string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return or
	}
	return value
}

func NewApiConfig() ApiConfig {
	if err := godotenv.Load(); err != nil {
		logger.Log.Error("ENV file not loaded", slog.Any("error", err))
	}

	return ApiConfig{
		Driver: tryGetEnv("DRIVER", storage.MemDriverName),
		Storage: StorageConfig{
			ConnectionString: tryGetEnv("CONNECTION", ""),
			Database:         tryGetEnv("DATABASE", "main"),
		},
	}
}
