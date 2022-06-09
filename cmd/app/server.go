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
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/incident"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var (
	grpcPort     = flag.Int("grpcPort", 9200, "gRPC por")
	grpcServer   *grpc.Server
	fleet_csv    = "../data/lfb_fleet_list_oct_2019.csv"
	incident_csv = "../data/animal_rescue_incidents_attended_lfb_from_jan_2009.csv"
)

func serve() {
	flag.Parse()

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
			}).WithError(err).Fatal()
		}

		fleetRepo := repository.NewRepository(fleet_csv)
		incidentRepo := repository.NewRepository(incident_csv)

		s1 := service.FleetServer{
			Repository:                      fleetRepo,
			UnimplementedFleetServiceServer: fleet.UnimplementedFleetServiceServer{},
		}

		s2 := service.IncidentServer{
			Repository:                         incidentRepo,
			UnimplementedIncidentServiceServer: incident.UnimplementedIncidentServiceServer{},
		}

		grpcServer := grpc.NewServer(
			grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
		)

		fleet.RegisterFleetServiceServer(grpcServer, &s1)
		incident.RegisterIncidentServiceServer(grpcServer, &s2)

		reflection.Register(grpcServer)

		log.WithFields(log.Fields{
			"message": "gRPC server: started...",
		}).Info(err)
		return grpcServer.Serve(listener)
	})

	select {
	case <-interrupt:
		log.WithFields(log.Fields{
			"msg": "Interrupt signal received...",
		}).WithContext(ctx).Warn()
		break
	case <-ctx.Done():
		log.WithFields(log.Fields{
			"msg": "Context done...",
		}).WithContext(ctx).Warn()
		break
	}

	log.WithFields(log.Fields{
		"message": "gRPC server: received shutdown signal...",
	}).Info()

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
		}).WithError(err).Warn()
		os.Exit(2)
	}
}
