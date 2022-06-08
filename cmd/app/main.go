package main

import (
	"net"
	"os"

	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/internal/service"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	fleet_csv = "../data/lfb_fleet_list_oct_2019.csv"
)

func main() {
	// Add logrus settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	repo := repository.NewRepository(fleet_csv)
	s1 := service.FleetServer{
		Repository:                      repo,
		UnimplementedFleetServiceServer: fleet.UnimplementedFleetServiceServer{},
	}

	// Configure the server
	listener, err := net.Listen("tcp", ":8086")

	if err != nil {
		log.WithError(err)
	}

	server := grpc.NewServer()
	fleet.RegisterFleetServiceServer(server, &s1)

	reflection.Register(server)

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
