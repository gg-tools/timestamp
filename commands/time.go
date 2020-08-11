package commands

import (
	"github.com/urfave/cli/v2"
	"time"
)

var Time = &cli.Command{
	Name:    "time",
	Aliases: []string{"t"},
	Usage:   "covert timestamp to time",
	UsageText: `timestamp t [timestamps...]

timestamp t 587433600
timestamp t 587433600 1597276800
`,
	ArgsUsage: "time [timestamp...]",
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			printTime(time.Now())
			return nil
		}

		for _, arg := range c.Args().Slice() {
			t, err := parseTimestamp(arg)
			if err != nil {
				return err
			}

			printTime(t)
		}

		return nil
	},
}
