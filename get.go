package debtool

import (
	"github.com/codegangsta/cli"
)

// RegisterGetCommand adds the "get" command to the command line interface.
func RegisterGetCommand(app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    "get",
		Aliases: []string{"g"},
		Usage:   "download packages from a repository",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "mirror, m",
				Value:  "http://archive.ubuntu.com/ubuntu",
				Usage:  "use the specified mirror",
				EnvVar: "DEBTOOL_MIRROR",
			},
			cli.StringFlag{
				Name:   "dist, d",
				Value:  "xenial",
				Usage:  "get packages from specified distribution",
				EnvVar: "DEBTOOL_DISTRIBUTION",
			},
		},
	})
}
