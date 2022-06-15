package service

import (
	"context"
	"github.com/shopspring/decimal"
	"strings"

	"github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *FleetServer) GetFleetStats(ctx context.Context, req *fleet.GetFleetStatsRequest) (*fleet.GetFleetStatsResponse, error) {
	fleetHandler, err := NewFleetHandler(s.Repository)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	yearFilter := repository.Filter{
		Property: "Reg Year",
		Value:    "0",
	}

	makeFilter := repository.Filter{
		Property: "Make",
		Value:    "",
	}

	var filters = make([]*repository.Filter, 0)

	if req.GetYear() != 0 {
		year := decimal.NewFromInt32(req.GetYear())
		yearFilter.Value = year.StringFixed(0)
		filters = append(filters, &yearFilter)
	}

	if len(req.GetMake()) > 0 {
		makeFilter.Value = req.GetMake()
		filters = append(filters, &makeFilter)
	}

	stats, err := fleetHandler.GetStats(filters)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &fleet.GetFleetStatsResponse{
		Active:     stats.Active,
		Reserve:    stats.Reserve,
		Training:   stats.Training,
		AverageAge: stats.AverageAge,
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
