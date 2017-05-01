package main

import (
    "os"

    "github.com/urfave/cli"
    "github.com/agile6v/gang/agent"
    "github.com/agile6v/gang/center"
)

func main() {
    app := cli.NewApp()
    app.Name = "Task Management System"
    app.Usage = "Usage oooooooo!"

    app.Commands = []cli.Command{
        center.Command(),
        agent.Command(),
    }

    app.Run(os.Args)
}
