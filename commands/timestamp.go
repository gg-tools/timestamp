package commands

import (
	"github.com/urfave/cli/v2"
	"time"
)

var Timestamp = &cli.Command{
	Name:    "timestamp",
	Aliases: []string{"ts"},
	Usage:   "covert time to timestamp",
	UsageText: `timestamp ts [date|datetime|expressions...]

timestamp ts 
timestamp ts 1988-08-13
timestamp ts "1988-08-13 18:06:06"
timestamp ts +1y -2m 3d 4h 5i 6s
`,
	ArgsUsage: `timestamp [date|datetime|expression...]`,
	Action: func(c *cli.Context) error {
		nArg := c.NArg()
		args := c.Args().Slice()
		if nArg == 0 {
			printTimestamp(time.Now())
			return nil
		} else if nArg == 1 && (len(args[0]) == dateLen || len(args[0]) == timeLen) {
			t, err := parseTime(args[0])
			if err != nil {
				return err
			}
			printTimestamp(t)
			return nil
		} else {
			t, err := str2time(args...)
			if err != nil {
				return err
			}
			printTimestamp(t)
			return nil
		}
	},
}
