package service

import (
	"context"
	"strings"

	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
	log "github.com/sirupsen/logrus"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/apimachinery/pkg/util/rand"
)

type FleetServer struct {
	fleet.UnimplementedFleetServiceServer
	Repository *repository.Repository
}

func (s *FleetServer) ListVehicles(ctx context.Context, in *emptypb.Empty) (*fleet.VehicleList, error) {
	fleetHandler, err := NewFleetHandler(s.Repository)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	vehicles, err := fleetHandler.GetVehicles()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &fleet.VehicleList{
		Vehicles: vehicles,
	}, nil
}

func (s *FleetServer) GetVehiclesByOpStatus(ctx context.Context, req *fleet.GetVehiclesByOpStatusRequest) (*fleet.VehicleList, error) {
	fleetHandler, err := NewFleetHandler(s.Repository)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	vehicles, err := fleetHandler.GetVehiclesByOperationalStatus(req.GetStatus())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &fleet.VehicleList{
		Vehicles: vehicles,
	}, nil
}

func (s *FleetServer) GetVehiclesByYear(ctx context.Context, req *fleet.GetVehiclesByYearRequest) (*fleet.VehicleList, error) {
	fleetHandler, err := NewFleetHandler(s.Repository)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	vehicles, err := fleetHandler.GetVehiclesByYear(req.GetYear())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &fleet.VehicleList{
		Vehicles: vehicles,
	}, nil
}

func (s *FleetServer) AddVehicle(ctx context.Context, req *fleet.VehicleRequest) (*fleet.VehicleResponse, error) {
	fleetHandler, err := NewFleetHandler(s.Repository)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	vehicle := &fleet.Vehicle{
		FleetNumber:       strings.ToUpper(rand.String(5)),
		OperationalStatus: req.GetOperationalStatus(),
		Lfb:               req.GetLfb(),
		Make:              req.GetMake(),
		Model:             req.GetModel(),
		Type:              req.GetType(),
		Category:          req.GetCategory(),
		RegistrationYear:  req.GetRegistrationYear(),
		Life:              req.GetLife(),
	}

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

func NewFleetServer(f string) (*FleetServer, error) {
	r := repository.NewRepository(f)

	return &FleetServer{
		Repository: r,
	}, nil
}
