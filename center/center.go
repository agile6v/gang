package center

import (
    "fmt"
    "github.com/urfave/cli"
)

func Command() cli.Command {
    return cli.Command{
        Name:   "center",
        Usage:  "start a center server",
        Flags:  []cli.Flag{
            cli.StringFlag{Name: "port", Value: "8315", Usage: "listening port"},
        },
        Action:  start,
    }
}

func start(c *cli.Context) error {
    fmt.Println("agent start")
}
