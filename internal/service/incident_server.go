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

func (s *IncidentServer) GetIncidentsByAnimalGroup(ctx context.Context, req *incident.GetIncidentsByAnimalGroupRequest) (*incident.IncidentList, error) {
	handler, err := NewIncidentHandler(s.Repository)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	incidents, err := handler.GetIncidentsByAnimalGroup(req.GetGroup())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &incident.IncidentList{
		Incidents: incidents,
	}, nil
}

func (s *IncidentServer) GetYearStats(ctx context.Context, req *incident.GetIncidentsStatsRequest) (*incident.GetIncidentsStatsResponse, error) {
	handler, err := NewIncidentHandler(s.Repository)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stats, err := handler.GetYearStats(req.GetYear())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	groupStats := make([]*incident.GetIncidentsStatsResponse_GroupStats, 0)

	for _, gsx := range stats.groups {
		gs := incident.GetIncidentsStatsResponse_GroupStats{
			AnimalGroup: gsx.animal_group,
			Incidents:   gsx.incidents,
		}

		groupStats = append(groupStats, &gs)
	}

	return &incident.GetIncidentsStatsResponse{
		Year:      stats.year,
		Incidents: stats.incidents,
		Groups:    groupStats,
	}, nil
}

func NewIncidentServer(f string) (*IncidentServer, error) {
	r := repository.NewRepository(f)

	return &IncidentServer{
		Repository: r,
	}, nil
}
