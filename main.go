package main

import (
	"github.com/gg-tools/timestamp/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "Timestamp"
	app.Compiled = time.Now()
	app.Usage = "Timestamp - conversion between time and timestamp"
	app.UsageText = `timestamp [command] [args...]

timestamp 
timestamp 1988-08-13
timestamp "1988-08-13 18:06:06"
timestamp +1y -2m 3d 4h 5i 6s
`
	app.Commands = []*cli.Command{commands.Time, commands.Timestamp}
	app.Action = commands.Timestamp.Action
	app.Flags = commands.Timestamp.Flags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
