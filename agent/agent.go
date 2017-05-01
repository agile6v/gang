package agent

import (
    "fmt"
    "github.com/urfave/cli"
)

func Command() cli.Command {
    return cli.Command{
        Name:   "agent",
        Usage:  "start a agent server",
        Flags:  []cli.Flag{
            cli.StringFlag{Name: "center", Value: "", Usage: "a list of center addresses, format likes 192.168.1.13:9898"},
            cli.StringFlag{Name: "port", Value: "8312", Usage: "listening port"},
        },
        Action: func(c *cli.Context) {
            center := c.String("center")
            fmt.Println("agent start", center)
        },
    }
}
