package main

import (
    "os"
    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "Task Management System"
    app.Usage = "Usage oooooooo!"

    app.Commands = []cli.Command{
        CenterCommand(),
        AgentCommand(),
    }

    app.Run(os.Args)
}
