package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type AgentCli struct {
	*cli.App
	// cache storage.FileStorage
}

func (c *AgentCli) PrintAgent(text string) {
	fmt.Printf("[Agent]: %s\n", text)
}

func (c *AgentCli) PrintServer(text string) {
	fmt.Printf("[Server]: %s\n", text)
}

func (c *AgentCli) Exit() {
	c.PrintAgent("Exit")
	os.Exit(0)
}

func (c *AgentCli) Waiting() {
	c.PrintAgent("Waiting...")
}

func NewApp() *AgentCli {
	return &AgentCli{
		&cli.App{
			Name:  "cliga",
			Usage: "CLI for LIGA agent",
		},
		// storage, // TODO
	}
}

func (c *AgentCli) Init() {

	var subSprint = cli.Command{
		Name:    "sprint",
		Aliases: []string{"s"},
		Usage:   "Get info of sprint by [number]",
		Args:    true,
		Action: func(ctx *cli.Context) error {
			sprintNum := ctx.Args().First()
			if sprintNum == "" {
				c.PrintAgent("The sprint number should not be empty")
				c.Exit()
				return nil
			}
			c.Waiting()
			// TODO

			c.Exit()
			return nil
		},
	}

	var subUser = cli.Command{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "Get user sprints by [uuid]",
		Args:    true,
		Action: func(ctx *cli.Context) error {

			uuid := ctx.Args().First()
			if uuid == "" {
				c.PrintAgent("The uuid should not be empty")
				c.Exit()
				return nil
			}
			c.Waiting()
			// TODO

			c.Exit()
			return nil
		},
	}

	var getCommand = cli.Command{
		Name:  "get",
		Usage: "Get [payload] from server",
		Subcommands: []*cli.Command{
			&subSprint, &subUser,
		},
	}

	var checkCommand = cli.Command{
		Name:  "check",
		Usage: "Check [number sprint]",
		Args:  true,
		Action: func(ctx *cli.Context) error {
			sprintNum := ctx.Args().First()
			if sprintNum == "" {
				c.PrintAgent("The sprint number should not be empty")
				c.Exit()
				return nil
			}
			c.Waiting()
			// TODO

			c.Exit()
			return nil
		},
	}
	c.Commands = append(c.Commands, &getCommand, &checkCommand)
}
