package cli

import (
	"github.com/urfave/cli/v2"
)

type AgentCli struct {
	*cli.App
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
			Usage:    "Concatinate first letter of firstname and full lastname. ex: iivanov",
			Required: true,
		}),
		Action: func(ctx *cli.Context) error {
			sprintNum := ctx.Args().First()
			if sprintNum == "" {
				printAgentError("The sprint number should not be empty", nil, ctx.Bool("debug"))
				printExit()
				return nil
			}
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
