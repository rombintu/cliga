package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlc "net/url"
	"os"

	"github.com/rombintu/checker-sprints/internal/storage"
	"github.com/urfave/cli/v2"
)

const (
	jsonHeader    string = "application/json"
	machineIDPath string = "/etc/machine-id"
)

type Url struct {
	path string
}

func NewUrl(base string) *Url {
	return &Url{path: base}
}

func (u *Url) addRoute(route string) {
	u.path = u.path + route
}

func (u *Url) addParams(params map[string]string) {
	query := urlc.Values{}
	for k, v := range params {
		query.Add(k, v)
	}
	u.path = fmt.Sprintf("%s?%s", u.path, query.Encode())
}

func (u *Url) addQueryParam(param string) {
	u.path = u.path + "/" + param
}

func (u *Url) Get() (string, error) {
	resp, err := http.Get(u.path)
	if err != nil {
		return "No connect to server", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "No response from server", err
	}
	return string(body), nil
}

func (u *Url) Post(payload storage.User) (string, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return "Marshal payload error", err
	}
	resp, err := http.Post(u.path, jsonHeader, &buf)
	if err != nil {
		return "No connect to server", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "No response from server", err
	}
	return string(body), nil
}

func getMachineID() (string, error) {
	// Проверяем, существует ли файл /etc/machine-id
	if _, err := os.Stat(machineIDPath); os.IsNotExist(err) {
		return "", err
	}

	// Читаем содержимое файла
	machineIDBytes, err := os.ReadFile(machineIDPath)
	if err != nil {
		return "", err
	}

	// Убираем символы новой строки
	machineID := string(machineIDBytes)
	machineID = machineID[:len(machineID)-1]

	return machineID, nil
}

func (c *AgentCli) ActionPing(ctx *cli.Context) {
	printAgent("PING!")
	url := NewUrl(ctx.String("server"))
	a, err := url.Get()
	if err != nil {
		printServerError(a, err, ctx.Bool("debug"))
	}
	printServer(a)
}

func (c *AgentCli) ActionSprintGet(ctx *cli.Context, sprintNum string) {
	switch sprintNum {
	case "1", "one", "first", "vpn":
		printSprint(sprintVPN)
	case "2", "two", "second", "fs":
		printSprint(sprintFS)
	default:
		printAgentError(fmt.Sprintf("Sprint [%s] not found", sprintNum), nil, false)
	}
}

func (c *AgentCli) ActionSprintCheck(ctx *cli.Context, sprintNum string) {
	debug := ctx.Bool("debug")
	uniqObject, err := getMachineID()
	if err != nil {
		printAgentError("Not found machine-id", err, debug)
	}
	user := storage.User{
		Name:   ctx.String("user"),
		Anchor: uniqObject,
	}

	printAgent(prettyInfo(fmt.Sprintf("Username: %s [%s]", user.Name, user.Anchor)))

	url := NewUrl(ctx.String("server"))
	url.addRoute("/users")
	url.addQueryParam(sprintNum)
	a, err := url.Get()
	if err != nil {
		printServerError(a, err, ctx.Bool("debug"))
	}
	printServer(a)
}
