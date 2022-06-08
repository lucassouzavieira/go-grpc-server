package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/internal/service"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var (
	version    = "v1"
	grpcPort   = flag.Int("grpcPort", 9200, "gRPC por")
	grpcServer *grpc.Server
	fleet_csv  = "../data/lfb_fleet_list_oct_2019.csv"
)

func serve() {
	flag.Parse()

	// Add logrus settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	// Interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	defer signal.Stop(interrupt)

	// gRPC server
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))

		if err != nil {
			log.WithFields(log.Fields{
				"error": "gRPC server: failed to listen",
			}).WithError(err)
		}

		repo := repository.NewRepository(fleet_csv)
		s1 := service.FleetServer{
			Repository:                      repo,
			UnimplementedFleetServiceServer: fleet.UnimplementedFleetServiceServer{},
		}

		grpcServer := grpc.NewServer(
			grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
		)

		fleet.RegisterFleetServiceServer(grpcServer, &s1)
		reflection.Register(grpcServer)

		log.WithFields(log.Fields{
			"message": "gRPC server: started...",
		}).WithError(err)
		return grpcServer.Serve(listener)
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.WithFields(log.Fields{
		"message": "gRPC server: received shutdown signal...",
	})

	cancel()

	_, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	err := g.Wait()

	if err != nil {
		log.WithFields(log.Fields{
			"message": "Server returning an error...",
		}).WithError(err)
		os.Exit(2)
	}
}
