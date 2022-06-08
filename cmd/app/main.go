package main

import (
	"net"
	"os"

	log "github.com/sirupsen/logrus"

	igrpc "github.com/lucassouzavieira/go-grpc-server/internal/grpc"
	pbf "github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/schema/fleet"
	"google.golang.org/grpc"
)

var (
// fleet = "../data/lfb_fleet_list_oct_2019.csv"
)

func main() {
	// Add logrus settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	s1 := igrpc.FleetServer{}

	// Configure the server
	listener, err := net.Listen("tcp", ":8086")

	if err != nil {
		log.WithError(err)
	}

	server := grpc.NewServer()
	pbf.RegisterFleetServiceServer(server, &s1)

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
