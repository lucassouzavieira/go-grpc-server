package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/incident"
)

func newConnection(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	if addr == "" {
		logrus.Fatalf("No address added.")
		return nil, nil
	}

	// Connection without TLS
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		logrus.WithError(err).Fatalf("fail to dial: %v")
		return nil, err
	}

	return conn, nil
}

func newFleetClient(addr string) (fleet.FleetServiceClient, error) {
	conn, err := newConnection(context.Background(), addr)

	if err != nil {
		logrus.WithError(err).Fatalf("fail to get grpc connection: %v")
		return nil, err
	}

	return fleet.NewFleetServiceClient(conn), nil
}

func newIncidentClient(addr string) (incident.IncidentServiceClient, error) {
	conn, err := newConnection(context.Background(), addr)

	if err != nil {
		logrus.WithError(err).Fatalf("fail to get grpc connection: %v")
		return nil, err
	}

	return incident.NewIncidentServiceClient(conn), nil
}
