package main

import (
	"github.com/gg-tools/timestamp/commands"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "Timestamp"
	app.Compiled = time.Now()
	app.Usage = "Timestamp - conversion between time and timestamp"
	app.UsageText = `timestamp [command] [args...]`
	app.Commands = []cli.Command{commands.Time, commands.Timestamp}
	app.Action = commands.Timestamp.Action

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
