package main

import (
	"log/slog"
	"os"

	"github.com/rombintu/checker-sprints/internal/cli"
)

func main() {
	c := cli.NewApp()
	c.Init()
	if err := c.Run(os.Args); err != nil {
		slog.Error(err.Error())
	}
}
