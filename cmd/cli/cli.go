package main

import (
	"context"
	"fmt"
	"time"

	cli "github.com/jawher/mow.cli"
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
		incidentsClient, err := newIncidentClient(*server)

		if err != nil {
			logrus.WithError(err).Fatal("Failed to initialize grpc Client...")
		}

		cmd.Command("list", "List incidents", func(cmd *cli.Cmd) {
			var (
				group = cmd.StringOpt("from-group", "", "List all incidents by animal group")
			)

			cmd.Action = func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

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
		// list --all
		// list --op-status
		// list --year
		// add
		fleetClient, err := newFleetClient(*server)

		if err != nil {
			logrus.WithError(err).Fatal("Failed to initialize grpc Client...")
		}

		cmd.Command("list", "List vehicles", func(cmd *cli.Cmd) {
			var (
				all    = cmd.BoolOpt("all", false, "List all vehicles")
				status = cmd.StringOpt("op-status", "", "List all vehicles by operational status")
				year   = cmd.IntOpt("year", 0, "List all vehicles by year")
			)

			cmd.Action = func() {
				fmt.Printf("List all vehicles? %v", *all)
				fmt.Printf("List by op-status? %v", *status)
				fmt.Printf("List by year? %v", *year)

				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				// List all
				if *all {
					resp, err := fleetClient.ListVehicles(ctx, &emptypb.Empty{})

					if err != nil {
						logrus.WithError(err).Fatal("Failed to get vehicle list")
					}

					fmt.Println(resp)
					return
				}
			}
		})

		cmd.Command("add", "Add a new vehicle", func(cmd *cli.Cmd) {
			// TODO implement
			cmd.Action = func() {
				logrus.Info("Add a new vehicle")
			}
		})
	})

	return app
}
