package commands

import (
	"github.com/urfave/cli"
	"time"
)

var Time = cli.Command{
	Name:      "time",
	ShortName: "t",
	Usage:     "covert timestamp to time",
	UsageText: `
timestamp time 587433600
timestamp time 587433600 1597276800
`,
	ArgsUsage: "time [timestamp...]",
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			printTime(time.Now())
			return nil
		}

		for _, arg := range c.Args() {
			t, err := parseTimestamp(arg)
			if err != nil {
				return err
			}

			printTime(t)
		}

		return nil
	},
}