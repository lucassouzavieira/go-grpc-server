package main

import (
	"context"
	"fmt"
	"time"

	cli "github.com/jawher/mow.cli"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/incident"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newCliApplication() *cli.Cli {
	app := cli.App("grpc-client", "A gRPC CLI client for go-grpc-server project")
	app.Spec = "[-s]"

	var (
		server = app.StringOpt("s server", "localhost:8080", "gRPC server address")
	)

	app.Before = func() {
		if *server == "" {
			logrus.Panic("Server address not given")
		}
	}

	// Incidents commands
	app.Command("incidents", "Handle LFB incidents info", func(cmd *cli.Cmd) {
		conn, err := newConnection(context.Background(), *server)

		if err != nil {
			logrus.WithError(err).Fatalf("fail to get grpc connection: %s", *server)
			return
		}

		incidentsClient := incident.NewIncidentServiceClient(conn)

		cmd.Command("list", "List incidents", func(cmd *cli.Cmd) {
			var (
				group = cmd.StringOpt("from-group", "", "List all incidents by animal group")
			)

			cmd.Action = func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				defer conn.Close()

				// List from group
				if len(*group) > 0 {
					resp, err := incidentsClient.GetIncidentsByAnimalGroup(ctx, &incident.GetIncidentsByAnimalGroupRequest{Group: *group})

					if err != nil {
						logrus.WithError(err).Fatal("Failed to get incidents from animal group")
					}

					fmt.Println(resp)
					return
				}

				// List all
				resp, err := incidentsClient.ListIncidents(ctx, &emptypb.Empty{})

				if err != nil {
					logrus.WithError(err).Fatal("Failed to get incidents")
				}

				fmt.Println(resp)
			}
		})
	})

	app.Command("fleet", "Handle LFB fleet info", func(cmd *cli.Cmd) {
		conn, err := newConnection(context.Background(), *server)

		if err != nil {
			logrus.WithError(err).Fatalf("fail to get grpc connection: %s", *server)
			return
		}

		fleetClient := fleet.NewFleetServiceClient(conn)

		cmd.Command("list", "List vehicles", func(cmd *cli.Cmd) {
			var (
				all    = cmd.BoolOpt("all", false, "List all vehicles")
				status = cmd.StringOpt("op-status", "", "List all vehicles by operational status")
				year   = cmd.IntOpt("year", 0, "List all vehicles by year")
			)

			cmd.Action = func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				defer conn.Close()

				// List all
				if *all {
					resp, err := fleetClient.ListVehicles(ctx, &emptypb.Empty{})

					if err != nil {
						logrus.WithError(err).Fatal("Failed to get vehicle list")
					}

					fmt.Println(resp)
					return
				}

				// Get by status
				if len(*status) > 0 {
					resp, err := fleetClient.GetVehiclesByOpStatus(ctx, &fleet.GetVehiclesByOpStatusRequest{
						Status: *status,
					})

					if err != nil {
						logrus.WithError(err).Fatal("Failed to get vehicle list by status")
					}

					fmt.Println(resp)
					return
				}

				// Get by Year
				if *year > 0 {
					resp, err := fleetClient.GetVehiclesByYear(ctx, &fleet.GetVehiclesByYearRequest{
						Year: int32(*year),
					})

					if err != nil {
						logrus.WithError(err).Fatal("Failed to get vehicle list by year")
					}

					fmt.Println(resp)
					return
				}
			}
		})

		cmd.Command("add", "Add a new vehicle", func(cmd *cli.Cmd) {
			var (
				lfb      = cmd.StringOpt("lfb", "LFB", "LFB")
				model    = cmd.StringArg("MODEL", "", "Vehicle model")
				make     = cmd.StringArg("MAKE", "", "Vehicle maker")
				vType    = cmd.StringOpt("type", "Car", "Vehicle type")
				category = cmd.StringOpt("category", "N3", "Vehicle category")
				status   = cmd.StringOpt("op-status", "ACTIVE", "Vehicle operational status")
				year     = cmd.IntOpt("year", time.Now().Year(), "Registration year")
				life     = cmd.IntOpt("life", 8, "Vehicle estimated lifetime")
			)

			cmd.Action = func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				defer conn.Close()

				// Add a new vehicle
				resp, err := fleetClient.AddVehicle(ctx, &fleet.VehicleRequest{
					OperationalStatus: *status,
					Lfb:               *lfb,
					Make:              *make,
					Model:             *model,
					Type:              *vType,
					Category:          *category,
					RegistrationYear:  int32(*year),
					Life:              int32(*life),
				})

				if err != nil {
					logrus.WithError(err).Fatal("Failed to add new vehicle list")
				}

				fmt.Println(resp)
			}
		})
	})

	return app
}
