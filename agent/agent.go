package agent

import (
    "fmt"
    "errors"
    "github.com/urfave/cli"
    . "github.com/agile6v/gang/util"
)

type Agent struct {
    listenIp string
    listenPort int
}

func Command() cli.Command {
    return cli.Command{
        Name:   "agent",
        Usage:  "start a agent server",
        Flags:  []cli.Flag{
            cli.StringFlag{Name: "center", Value: "", Usage: "a list of center addresses, format likes 192.168.1.13:9898"},
            cli.StringFlag{Name: "ip", Value: "localhost", Usage: "listening IP address"},
            cli.StringFlag{Name: "port", Value: "8312", Usage: "listening port"},
        },
        Action: start,
    }
}

func start(c *cli.Context) error {
    center := c.String("center")
    if center == "" {
        ReturnError(EXIT_BAD_ARGS, errors.New("center is required"))
    }
    fmt.Println("agent start", center)
}

func init(ip string, port string, center string) *Agent {
    agent := &Agent{

    }
}

func (agent *Agent) run() {
    http.HandleFunc("/jobs", tl.statusHandler)
    http.ListenAndServe(agent.listenIp , nil)
}
