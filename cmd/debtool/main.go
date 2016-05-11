package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/strothj/alpine-debtool"
)

func main() {
	app := cli.NewApp()
	debtool.RegisterGetCommand(app)
	app.Run(os.Args)
}
