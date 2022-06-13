package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		logrus.WithError(err).Fatalf("fail to dial: %s", addr)
		return nil, err
	}

	return conn, nil
}
