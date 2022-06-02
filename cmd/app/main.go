package main

import (
	"fmt"
	"log"
	"os"

	repository "github.com/lucassouzavieira/go-grpc-server/internal/repository"
)

var (
	mockup = "../test/data/mockup.csv"
)

func main() {
	logger := log.New(os.Stderr, "App info: ", 2)
	repo := repository.NewRepository(mockup)
	data, err := repo.GetData()

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Print(data)
}
