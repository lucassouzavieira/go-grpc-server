package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Add logrus settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	serve()
}
