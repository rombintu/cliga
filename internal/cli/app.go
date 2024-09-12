package cli

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"
)

const addressAPI = "http://192.168.213.204:8080"

type AgentCli struct {
	*cli.App
	// cache storage.FileStorage
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
)

func PrintAgent(text string) {
	fmt.Printf("["+ColorPurple+"Agent"+ColorReset+"]: %s\n", text)
}

func PrintServer(text string) {
	fmt.Printf("["+ColorGreen+"Server"+ColorReset+"]: %s\n", text)
}

func PrintExit() {
	PrintAgent("Exit")
	os.Exit(0)
}

func PrintWaiting() {
	PrintAgent("Waiting...")
}

func PrintServerError(text string, err error, debug bool) {
	PrintServer(ColorRed + text + ColorReset)
	if debug {
		slog.Error(err.Error())
	}
	PrintExit()
}

func PrintAgentError(text string, err error, debug bool) {
	PrintAgent(ColorRed + text + ColorReset)
	if debug {
		slog.Error(err.Error())
	}
	PrintExit()
}

func NewApp() *AgentCli {
	return &AgentCli{
		App: &cli.App{
			Name:  "cliga",
			Usage: "CLI for LIGA agent",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "debug",
					Value: false,
				},
			},
		},
		// storage, // TODO
	}
}

func (c *AgentCli) Init() {
	var flagServer = &cli.StringFlag{
		Name:  "server",
		Value: addressAPI,
	}
	var defaultFlagsForServer = []cli.Flag{
		flagServer,
	}

	var subSprint = cli.Command{
		Name:    "sprint",
		Aliases: []string{"s"},
		Usage:   "Get info of sprint by [number]",
		Args:    true,
		Flags:   defaultFlagsForServer,
		Action: func(ctx *cli.Context) error {
			sprintNum := ctx.Args().First()
			if sprintNum == "" {
				PrintAgent("The sprint number should not be empty")
				PrintExit()
				return nil
			}
			PrintWaiting()
			// TODO
			c.ActionSprintGet(ctx, sprintNum)
			PrintExit()
			return nil
		},
	}

	var subUser = cli.Command{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "Get user sprints by [uuid]",
		Args:    true,
		Flags:   defaultFlagsForServer,
		Action: func(ctx *cli.Context) error {

			uuid := ctx.Args().First()
			if uuid == "" {
				PrintAgent("The uuid should not be empty")
				PrintExit()
				return nil
			}
			PrintWaiting()
			// TODO

			PrintExit()
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
		Flags: defaultFlagsForServer,
		Action: func(ctx *cli.Context) error {
			sprintNum := ctx.Args().First()
			if sprintNum == "" {
				PrintAgentError("The sprint number should not be empty", nil, ctx.Bool("debug"))
				PrintExit()
				return nil
			}
			PrintWaiting()
			// TODO

			PrintExit()
			return nil
		},
	}

	var pingCommand = cli.Command{
		Name:  "ping",
		Usage: "Simple ping server",
		Flags: defaultFlagsForServer,
		Action: func(ctx *cli.Context) error {
			c.ActionPing(ctx)
			return nil
		},
	}
	c.Commands = append(c.Commands, &getCommand, &checkCommand, &pingCommand)
}
