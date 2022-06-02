package main

import (
	"fmt"
	"log"
	"os"

	grpc "github.com/lucassouzavieira/go-grpc-server/internal/grpc"
	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
)

var (
	fleet = "../data/lfb_fleet_list_oct_2019.csv"
)

func main() {
	logger := log.New(os.Stderr, "App info: ", 2)
	handler, err := grpc.NewFleetHandler(repository.NewRepository(fleet))

	if err != nil {
		logger.Fatal(err)
	}

	vehicles, err := handler.GetVehicles()

	if err != nil {
		logger.Fatal(err)
	}


	fmt.Print(vehicles[0].FleetNumber, vehicles[0].Lfb)
}
