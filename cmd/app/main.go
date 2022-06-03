package main

import (
	"log"
	"net"
	"os"

	igrpc "github.com/lucassouzavieira/go-grpc-server/internal/grpc"
	pbf "github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/schema/fleet"
	"google.golang.org/grpc"
)

var (
	// fleet = "../data/lfb_fleet_list_oct_2019.csv"
)

func main() {
	logger := log.New(os.Stderr, "App info: ", 2)
	s1 := igrpc.FleetServer{}

	// Configure the server
	listener, err := net.Listen("tcp", ":8086")

	if err != nil {
		logger.Fatal(err)
	}

	server := grpc.NewServer()
	pbf.RegisterFleetServiceServer(server, &s1)

	if err := server.Serve(listener); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
