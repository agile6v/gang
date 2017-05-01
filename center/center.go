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
        Action: func(c *cli.Context) {
            fmt.Println("center start")
        },
    }
}
