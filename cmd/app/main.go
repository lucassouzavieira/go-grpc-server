package main

import (
	web "github.com/lucassouzavieira/go-grpc-server/internal/http"
	logging "github.com/lucassouzavieira/go-grpc-server/internal/logging"
)

func main() {
	logging.LogMessage("Logging out...")
	web.InitHttpServer()
}
