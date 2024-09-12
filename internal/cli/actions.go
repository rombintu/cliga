package cli

import (
	"fmt"
	"io"
	"net/http"
	urlc "net/url"

	"github.com/urfave/cli/v2"
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

func (c *AgentCli) ActionPing(ctx *cli.Context) {
	PrintAgent("PING!")
	url := NewUrl(ctx.String("server"))
	a, err := url.Get()
	if err != nil {
		PrintServerError(a, err, ctx.Bool("debug"))
	}
	PrintServer(a)
}

func (c *AgentCli) ActionSprintGet(ctx *cli.Context, sprintNum string) {
	url := NewUrl(ctx.String("server"))
	url.addRoute("/sprints")
	url.addQueryParam(sprintNum)
	a, err := url.Get()
	if err != nil {
		PrintServerError(a, err, ctx.Bool("debug"))
	}
	PrintServer(a)
}
