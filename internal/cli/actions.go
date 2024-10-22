package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/rombintu/checker-sprints/internal/storage"
	"github.com/urfave/cli/v2"
)

const (
	jsonHeader    string = "application/json"
	machineIDPath string = "/etc/machine-id"
)

var errNone = errors.New("none")

type Url struct {
	path string
}

func NewUrl(base string) *Url {
	printWaiting()
	return &Url{path: base}
}

func (u *Url) addRoute(route string) {
	u.path = u.path + route
}

// func (u *Url) addParams(params map[string]string) {
// 	query := urlc.Values{}
// 	for k, v := range params {
// 		query.Add(k, v)
// 	}
// 	u.path = fmt.Sprintf("%s?%s", u.path, query.Encode())
// }

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

func (u *Url) Post(payload storage.User) (string, int, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return "Marshal payload error", 0, err
	}
	resp, err := http.Post(u.path, jsonHeader, &buf)
	if err != nil {
		return "No connect to server", 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "No response from server", 0, err
	}
	return string(body), resp.StatusCode, nil
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
		printSprint(SprintVPN)
	case "2", "two", "second", "fs":
		printSprint(SprintFS)
	case "3", "three":
		printSprint(SprintGrep)
	case "4":
		printSprint(SprintLVM)
	case "5":
		printSprint(SprintDeamon)
	case "6":
		printSprint(SprintVLAN)
	case "7":
		printSprint(SprintOps)
	default:
		printAgentError(fmt.Sprintf("Sprint [%s] not found", sprintNum), errNone, false)
	}
}

func (c *AgentCli) ActionSprintCheck(ctx *cli.Context, sprintNum string) {
	debug := ctx.Bool("debug")
	uniqObject, err := getMachineID()
	if err != nil {
		printAgentError("Not found machine-id", err, debug)
	}
	user := storage.User{
		Login:  ctx.String("user"),
		Anchor: uniqObject,
	}

	printAgent(prettyInfo(fmt.Sprintf("Username: %s [%s]", user.Login, user.Anchor)))
	// sch := storage.NewModelSprints()
	var s *Sprint
	switch sprintNum {
	case "1", "one", "first", "vpn":
		s = SprintVPN
	case "2", "two", "second", "fs":
		s = SprintFS
	case "3", "three":
		s = SprintGrep
	case "4":
		s = SprintLVM
	case "5":
		s = SprintDeamon
	case "6":
		s = SprintVLAN
	case "7":
		s = SprintOps
	default:
		printAgentError(fmt.Sprintf("Sprint [%s] not found", sprintNum), errNone, false)
	}

	stepsOK := true
	for _, step := range s.Steps {
		if !step.Check() {
			stepsOK = false
			printAgent(fmt.Sprintf("Taks [%d] not solved", step.ID))
		}
	}
	if !stepsOK {
		printAgentError("Some tasks are not solved, completion of the verification process", errNone, false)
	}
	printAgent(fmt.Sprintf("%sAll the tasks in the Sprint [%d] are solved!%s", ColorGreen, s.ID, ColorReset))
	url := NewUrl(ctx.String("server"))
	url.addRoute("/users/sprint")
	url.addQueryParam(strconv.FormatInt(s.ID, 10))
	a, code, err := url.Post(user)
	if err != nil {
		printServerError(a, err, ctx.Bool("debug"))
	} else if code > 200 {
		printServerError(a, errNone, false)
	}
	printServer(a)
}

func (c *AgentCli) ActionUserGet(ctx *cli.Context, username string) {
	url := NewUrl(ctx.String("server"))
	url.addRoute("/users")
	url.addQueryParam(username)
	payload, err := url.Get()
	if err != nil {
		printServerError(payload, err, false)
	}
	var user storage.User
	if err := json.Unmarshal([]byte(payload), &user); err != nil {
		printServerError(payload, err, false)
	}
	layout := "2006-01-02 15:04:05"
	for _, spr := range user.Sprints {
		formattedTime := spr.UpdatedAt.Format(layout)
		printServer(fmt.Sprintf("Sprint %d - %s | %s", spr.ID, formattedTime, OK))
	}
}
