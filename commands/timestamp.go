package commands

import (
	"github.com/urfave/cli"
	"time"
)

var Timestamp = cli.Command{
	Name:      "timestamp",
	ShortName: "ts",
	Usage:     "covert time to timestamp",
	UsageText: `
timestamp
timestamp 1988-08-13
timestamp "1988-08-13 18:06:06"
timestamp +1y -2m 3d 4h 5i 6s
`,
	ArgsUsage: `timestamp [date|datetime|expression...]`,
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			printTimestamp(time.Now())
			return nil
		} else if c.NArg() == 1 && (len(c.Args()[0]) == dateLen || len(c.Args()[0]) == timeLen) {
			t, err := parseTime(c.Args()[0])
			if err != nil {
				return err
			}
			printTimestamp(t)
			return nil
		} else {
			t, err := str2time(c.Args()...)
			if err != nil {
				return err
			}
			printTimestamp(t)
			return nil
		}
	},
}
