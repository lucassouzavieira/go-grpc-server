package service

import (
	"context"

	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/incident"
	log "github.com/sirupsen/logrus"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type IncidentServer struct {
	incident.UnimplementedIncidentServiceServer
	Repository *repository.Repository
}

func (s *IncidentServer) ListIncidents(ctx context.Context, in *emptypb.Empty) (*incident.IncidentList, error) {
	handler, err := NewIncidentHandler(s.Repository)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	incidents, err := handler.GetIncidents()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &incident.IncidentList{
		Incidents: incidents,
	}, nil
}

func NewIncidentServer(f string) (*IncidentServer, error) {
	r := repository.NewRepository(f)

	return &IncidentServer{
		Repository: r,
	}, nil
}
