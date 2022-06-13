package main

import (
	"fmt"

	cli "github.com/jawher/mow.cli"
	"github.com/sirupsen/logrus"
)

func newCliApplication() *cli.Cli {
	app := cli.App("grpc-client", "A gRPC CLI client for go-grpc-server project")

	app.Command("incidents", "Handle LFB incidents info", func(cmd *cli.Cmd) {
		// list --all
		// list --from-group
		cmd.Command("list", "List incidents", func(cmd *cli.Cmd) {
			var (
				all   = cmd.BoolOpt("all", true, "List all incidents")
				group = cmd.StringOpt("from-group", "", "List all incidents by animal group")
			)

			cmd.Action = func() {
				fmt.Printf("List all incidents? %v", *all)
				fmt.Printf("List by group? %v", *group)
			}
		})
	})

	app.Command("fleet", "Handle LFB fleet info", func(cmd *cli.Cmd) {
		// list --all
		// list --op-status
		// list --year
		// add
		cmd.Command("list", "List vehicles", func(cmd *cli.Cmd) {
			var (
				all    = cmd.BoolOpt("all", true, "List all vehicles")
				status = cmd.StringOpt("op-status", "", "List all vehicles by operational status")
				year   = cmd.IntOpt("year", 0, "List all vehicles by year")
			)

			cmd.Action = func() {
				fmt.Printf("List all vehicles? %v", *all)
				fmt.Printf("List by op-status? %v", *status)
				fmt.Printf("List by year? %v", *year)
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
