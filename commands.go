package debtool

import "github.com/codegangsta/cli"

// AddCommandsToCLI adds the available commands to the command line interface.
func AddCommandsToCLI(app *cli.App) {
	RegisterGetCommand(app)
}
