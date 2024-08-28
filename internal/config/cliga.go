package config

import (
	"flag"
)

const version string = "sprint1:0.1.0"

type CliConfig struct {
	ApiURL        string `json:"api_url"`
	UserName      string `json:"user_name"`
	Version       string `json:"version"`
	ActualVersion string `json:"actual_version"`
	Debug         bool   `json:"debug"`
}

func NewCliConfig() CliConfig {
	return CliConfig{
		Version: version,
	}
}

func (c *CliConfig) ParseFromFlags() {

	debug := flag.Bool("v", false, "Debug mode")
	user := flag.String("u", "", "User last name for identify")
	flag.Parse()

	c.Debug = *debug
	c.UserName = *user
}
