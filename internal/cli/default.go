package cli

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"
)

// Законстантил TODO засунуть в конфиг
const addressAPI = "http://192.168.213.204:8080"

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
)

func (a *AgentCli) ActionGetVersion(ctx *cli.Context) {
	printAgent(fmt.Sprintf("Stable: oct 21. %s", prettyTitle("7.0")))
	printServer(fmt.Sprintf("Stable: oct 14. %s", prettyTitle("2.1")))
}

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

func printAgentWarn(text string, err error, debug bool) {
	printAgent(ColorYellow + text + ColorReset)
	if debug {
		slog.Error(err.Error())
	}
}

func prettyTitle(text string) string {
	return ColorYellow + text + ColorReset
}

func prettyInfo(text string) string {
	return ColorCyan + text + ColorReset
}
