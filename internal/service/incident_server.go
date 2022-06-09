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
	repo := s.Repository
	handler, err := NewIncidentHandler(repo)

	if err != nil {
		log.Fatal(err)
	}

	incidents, err := handler.GetIncidents()

	if err != nil {
		log.Fatal(err)
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
