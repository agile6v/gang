package agent

import (
    "fmt"
    "strconv"
    "errors"
    "strings"
    "io"
    "net/http"
    "github.com/urfave/cli"
    . "github.com/agile6v/gang/util"
)

type Agent struct {
    listenOn string
    centers []string
}

func Command() cli.Command {
    return cli.Command{
        Name:   "agent",
        Usage:  "start a agent server",
        Flags:  []cli.Flag{
            cli.StringFlag{Name: "centers", Value: "", Usage: "a list of center addresses, format likes 192.168.1.13:9898"},
            cli.StringFlag{Name: "ip", Value: "localhost", Usage: "listening IP address"},
            cli.IntFlag{Name: "port", Value: 8923, Usage: "listening port"},
        },
        Action: start,
    }
}

func start(c *cli.Context) {
    ip := c.String("ip")
    port := c.Int("port")
    centers := strings.Split(c.String("centers"), ",")
    if centers[0] == "" {
        ReturnError(EXIT_BAD_ARGS, errors.New("center is required"))
    }

    fmt.Println("centers:", centers)
    fmt.Println("local_ip:", ip)
    fmt.Println("local_port:", port)

    agent := newAgent(ip + ":" + strconv.Itoa(port), centers)
    agent.run()
}

func newAgent(listenOn string, centers []string) *Agent {
    agent := &Agent{
        listenOn: listenOn,
        centers: centers,
    }
    return agent
}

func (agent *Agent) run() {
    http.HandleFunc("/jobs", agent.getAllJobs)
    http.ListenAndServe(agent.listenOn , nil)
}

func (agent *Agent) getAllJobs(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "All jobs!\n")
}

