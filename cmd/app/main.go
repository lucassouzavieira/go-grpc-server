package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Add logrus settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	serve()
}
