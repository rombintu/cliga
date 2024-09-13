package cli

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/rombintu/checker-sprints/internal/storage"
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

func printAgent(text string) {
	fmt.Printf("["+ColorPurple+"Agent"+ColorReset+"]: %s\n", text)
}

func printServer(text string) {
	fmt.Printf("["+ColorGreen+"Server"+ColorReset+"]: %s\n", text)
}

func printExit() {
	printAgent("Exit")
	os.Exit(0)
}

func printWaiting() {
	printAgent("Waiting...")
}

func printServerError(text string, err error, debug bool) {
	printServer(ColorRed + text + ColorReset)
	if debug {
		slog.Error(err.Error())
	}
	printExit()
}

func printAgentError(text string, err error, debug bool) {
	printAgent(ColorRed + text + ColorReset)
	if debug {
		slog.Error(err.Error())
	}
	printExit()
}

func prettyTitle(text string) string {
	return ColorBlue + text + ColorReset
}

func prettyInfo(text string) string {
	return ColorCyan + text + ColorReset
}

func printSprint(s *storage.Sprint) {
	printAgent(
		fmt.Sprintf("Sprint %d. %s", s.ID, prettyTitle(s.Title)))
	for _, step := range s.GetSteps() {
		if step.Body == "" {
			continue
		}
		printAgent(
			fmt.Sprintf("%d. %s", step.ID, step.Body))
	}
}

func NewApp() *AgentCli {
	SprintsInit()
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
				printAgent("The sprint number should not be empty")
				printExit()
				return nil
			}
			printWaiting()
			// TODO
			c.ActionSprintGet(ctx, sprintNum)
			printExit()
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
				printAgent("The uuid should not be empty")
				printExit()
				return nil
			}
			printWaiting()
			// TODO
			// c.ActionUserGet(ctx, uuid)
			printExit()
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
		Flags: append(defaultFlagsForServer, &cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u", "username"},
			Usage:    "Concatinate first letter of firstname and full lastname. Ex: iivanov",
			Required: true,
		}),
		Action: func(ctx *cli.Context) error {
			sprintNum := ctx.Args().First()
			if sprintNum == "" {
				printAgentError("The sprint number should not be empty", nil, ctx.Bool("debug"))
				printExit()
				return nil
			}
			printWaiting()
			// TODO
			c.ActionSprintCheck(ctx, sprintNum)
			printExit()
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
