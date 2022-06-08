package service

// Loads data from a csv entry and transform to a proto model struct
import (
	log "github.com/sirupsen/logrus"
	"strconv"

	repository "github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
)

type FleetHandler struct {
	r *repository.Repository
	v []*fleet.Vehicle
}

func NewFleetHandler(r *repository.Repository) (*FleetHandler, error) {
	data, err := r.GetData()

	if err != nil {
		log.Fatal(err)
	}

	var vehicles []*fleet.Vehicle
	for i, v := range data {
		if i == 0 {
			continue // skip headers
		}

		vehicle := vehicleFromCsv(v)
		vehicles = append(vehicles, &vehicle)
	}

	return &FleetHandler{
		r: r,
		v: vehicles,
	}, nil
}

func (h *FleetHandler) GetVehicles() ([]*fleet.Vehicle, error) {
	return h.v, nil
}

func (h *FleetHandler) GetVehiclesByOperationalStatus(status string) ([]*fleet.Vehicle, error) {
	var filtered []*fleet.Vehicle = make([]*fleet.Vehicle, 0)

	for _, v := range h.v {
		if v.OperationalStatus == status {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}

// package internal functions
func vehicleFromCsv(s []string) fleet.Vehicle {
	registrationYear, err := strconv.Atoi(s[7])

	if err != nil {
		log.Fatal(err)
	}

	life, err := strconv.Atoi(s[8])

	if err != nil {
		log.Fatal(err)
	}

	return fleet.Vehicle{
		FleetNumber:       s[0],
		OperationalStatus: s[1],
		Lfb:               s[2],
		Make:              s[3],
		Model:             s[4],
		Type:              s[5],
		Category:          s[6],
		RegistrationYear:  int32(registrationYear),
		Life:              int32(life),
	}
}
