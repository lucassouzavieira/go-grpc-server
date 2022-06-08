package service

import (
	"context"

	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
	log "github.com/sirupsen/logrus"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	fleetData = "../data/lfb_fleet_list_oct_2019.csv"
)

type FleetServer struct {
	fleet.UnimplementedFleetServiceServer
}

func (s *FleetServer) ListVehicles(ctx context.Context, in *emptypb.Empty) (*fleet.VehicleList, error) {
	repo := repository.NewRepository(fleetData)
	fleetHandler, err := NewFleetHandler(repo)

	if err != nil {
		log.Fatal(err)
	}

	vehicles, err := fleetHandler.GetVehicles()

	if err != nil {
		log.Fatal(err)
	}

	return &fleet.VehicleList{
		Vehicles: vehicles,
	}, nil
}

func (s *FleetServer) GetVehiclesByOpStatus(ctx context.Context, req *fleet.GetVehiclesByOpStatusRequest) (*fleet.VehicleList, error) {
	repo := repository.NewRepository(fleetData)
	fleetHandler, err := NewFleetHandler(repo)

	if err != nil {
		log.Fatal(err)
	}

	vehicles, err := fleetHandler.GetVehiclesByOperationalStatus(req.GetStatus())

	if err != nil {
		log.Fatal(err)
	}

	return &fleet.VehicleList{
		Vehicles: vehicles,
	}, nil
}

func (s *FleetServer) AddVehicle(ctx context.Context, in *fleet.VehicleRequest) (*fleet.VehicleResponse, error) {
	response := &fleet.VehicleResponse{
		Vehicle: &fleet.Vehicle{
			FleetNumber:       "",
			OperationalStatus: "",
			Lfb:               "",
			Make:              "",
			Model:             "",
			Type:              "",
			Category:          "",
			RegistrationYear:  0,
			Life:              0,
		},
		Created: false,
	}

	return response, nil
}

func NewFleetServer() (*FleetServer, error) {
	return &FleetServer{}, nil
}
