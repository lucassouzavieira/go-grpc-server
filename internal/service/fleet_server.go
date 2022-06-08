package service

import (
	"context"

	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
	log "github.com/sirupsen/logrus"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/apimachinery/pkg/util/rand"
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

func (s *FleetServer) AddVehicle(ctx context.Context, req *fleet.VehicleRequest) (*fleet.VehicleResponse, error) {
	repo := repository.NewRepository(fleetData)
	fleetHandler, err := NewFleetHandler(repo)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	vehicle := &fleet.Vehicle{
		FleetNumber:       rand.String(5),
		OperationalStatus: req.GetOperationalStatus(),
		Lfb:               req.GetLfb(),
		Make:              req.GetMake(),
		Model:             req.GetModel(),
		Type:              req.GetType(),
		Category:          req.GetCategory(),
		RegistrationYear:  req.GetRegistrationYear(),
		Life:              req.GetLife(),
	}

	// TODO Add data
	res, err := fleetHandler.SaveVehicle(vehicle)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &fleet.VehicleResponse{
		Vehicle: vehicle,
		Created: res,
	}, nil
}

func NewFleetServer() (*FleetServer, error) {
	return &FleetServer{}, nil
}
