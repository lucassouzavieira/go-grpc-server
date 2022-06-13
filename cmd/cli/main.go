package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	app := newCliApplication()

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Panic("Failed to run CLI app")
	}
}
