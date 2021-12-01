package app

import "github.com/urfave/cli"

// Generate will return the command line application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search for IPs and server names on internet"

	return app
}
