package main

import (
	"os"

	"github.com/SEEK-Jobs/dynamotools/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		cmd.BuildArchive(),
		cmd.BuildRestore(),
	}

	app.Run(os.Args)
}
