package grpc

import (
	"context"

	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
	pbfleet "github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/schema/fleet"
	log "github.com/sirupsen/logrus"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	fleetData = "../data/lfb_fleet_list_oct_2019.csv"
)

type FleetServer struct {
	pbfleet.UnimplementedFleetServiceServer
}

func (s *FleetServer) ListVehicles(ctx context.Context, in *emptypb.Empty) (*pbfleet.VehicleList, error) {
	repo := repository.NewRepository(fleetData)
	fleetHandler, err := NewFleetHandler(repo)

	if err != nil {
		log.Fatal(err)
	}

	vehicles, err := fleetHandler.GetVehicles()

	if err != nil {
		log.Fatal(err)
	}

	return &pbfleet.VehicleList{
		Vehicles: vehicles,
	}, nil
}

func (s *FleetServer) AddVehicle(ctx context.Context, in *pbfleet.VehicleRequest) (*pbfleet.VehicleResponse, error) {
	response := &pbfleet.VehicleResponse{
		Vehicle: &pbfleet.Vehicle{
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
