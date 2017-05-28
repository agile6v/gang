package agent

import (
	"errors"
	"fmt"
	"io"
    "io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
    "sync"
    "bytes"
    "encoding/json"
	. "github.com/agile6v/gang/common"
	. "github.com/agile6v/gang/util"
	"github.com/urfave/cli"
)

type Agent struct {
	listenOn string
	centers  []string
	Tasks    map[int64]*Task
	Lock     sync.Mutex
}

func Command() cli.Command {
	return cli.Command{
		Name:  "agent",
		Usage: "start a agent server",
		Flags: []cli.Flag{
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

	agent := newAgent(ip+":"+strconv.Itoa(port), centers)
	go agent.heartbeat()
	agent.run()
}

func newAgent(listenOn string, centers []string) *Agent {
	agent := &Agent{
		listenOn: listenOn,
		centers:  centers,
	}
	return agent
}

func (agent *Agent) run() {
	http.HandleFunc("/jobs", agent.getAllJobs)
	http.ListenAndServe(agent.listenOn, nil)
}

func (agent *Agent) heartbeat() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("heartbeating...")
            err := agent.syncTasks()
            if err != nil {
			    fmt.Println("syncTask failed")
            }
		}
	}
	fmt.Println("heartbeat() quit")
}

func (agent *Agent) syncTasks() error {
	json, err := json.Marshal(agent.Tasks)
	if err != nil {
	    fmt.Println("syncTasks() json failed")
		return err
	}

    url := "http://" + agent.centers[0] + "/heartbeat"
    fmt.Println("-----------", bytes.NewBuffer(json).String())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

    return nil
}

func (agent *Agent) getAllJobs(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All jobs!\n")
}
