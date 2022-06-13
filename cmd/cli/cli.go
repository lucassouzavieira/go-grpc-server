package main

import (
	cli "github.com/jawher/mow.cli"
	"github.com/sirupsen/logrus"
)

func newCliApplication() *cli.Cli {
	app := cli.App("cli", "CLI App")

	app.Command("hello-world", "a simple cli command", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			logrus.Info("Hello World!")
		}
	})

	return app
}
