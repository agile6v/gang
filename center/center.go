package center

import (
    "fmt"
    "strconv"
    "net/http"
    "io"
    "github.com/urfave/cli"
)

type Center struct {
    listenOn string
}

func Command() cli.Command {
    return cli.Command{
        Name:   "center",
        Usage:  "start a center server",
        Flags:  []cli.Flag{
            cli.StringFlag{Name: "ip", Value: "localhost", Usage: "listening IP address"},
            cli.IntFlag{Name: "port", Value: 8902, Usage: "listening port"},
        },
        Action:  start,
    }
}

func start(c *cli.Context) {
    ip := c.String("ip")
    port := c.Int("port")

    fmt.Println("local_ip:", ip)
    fmt.Println("local_port:", port)

    center := newCenter(ip + ":" + strconv.Itoa(port))
    center.run()
}

func newCenter(listenOn string) *Center {
    center := &Center{
        listenOn: listenOn,
    }
    return center
}

func (center *Center) run() {
    http.HandleFunc("/heartbeat", center.handleHeartBeat)
    http.ListenAndServe(center.listenOn , nil)
}

func (center *Center) handleHeartBeat(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "HeartBeat!\n")
}
