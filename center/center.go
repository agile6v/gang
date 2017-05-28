package center

import (
	"fmt"
	. "github.com/agile6v/gang/common"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Center struct {
	listenOn string
	dsn      string
	TaskMap  map[string][]*Task
	Lock     sync.Mutex
}

func Command() cli.Command {
	return cli.Command{
		Name:  "center",
		Usage: "start a center server",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "ip", Value: "localhost", Usage: "listening IP address"},
			cli.IntFlag{Name: "port", Value: 8902, Usage: "listening port"},
			cli.StringFlag{Name: "db", Value: "root:@tcp(127.0.0.1:3306)/gang", Usage: "DSN format"},
		},
		Action: start,
	}
}

func start(c *cli.Context) {
	ip := c.String("ip")
	port := c.Int("port")
	dsn := c.String("db")

	fmt.Println("local_ip:", ip)
	fmt.Println("local_port:", port)
	fmt.Println("database:", dsn)

	center := newCenter(ip+":"+strconv.Itoa(port), dsn)
	db := &DB{Dsn: dsn}

	ret := center.getAllTasks(db)
	if !ret {
		os.Exit(-1)
	}

	go center.syncTasks(db)

	center.run()
}

func newCenter(listenOn string, database string) *Center {
	center := &Center{
		listenOn: listenOn,
		dsn:      database,
	}
	return center
}

func (center *Center) getAllTasks(db *DB) bool {
	m, err := db.GetAllTasks()
	if err != nil {
		fmt.Printf("GetAllTasks failure: ", err)
		return false
	}

	center.Lock.Lock()
	center.TaskMap = m
	center.Lock.Unlock()

	for host, tasks := range m {
		for _, task := range tasks {
			fmt.Printf("host=%v, task=%s,%d,%s,%s,%s\n", host, task.Name, task.Id, task.Command, task.Args, task.Runner)
		}
	}

	return true
}

func (center *Center) syncTasks(db *DB) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("center->getAllTasks()")
			center.getAllTasks(db)
		}
	}
	fmt.Println("center->getAllTasks() quit")
}

func (center *Center) run() {
	http.HandleFunc("/heartbeat", center.handleHeartBeat)
	http.ListenAndServe(center.listenOn, nil)
}

func (center *Center) handleHeartBeat(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	io.WriteString(w, string(body))
}
