package grpc

// Loads data from a csv entry and transform to a proto model struct
import (
	"log"
	"os"
	"strconv"

	repository "github.com/lucassouzavieira/go-grpc-server/internal/repository"
	pbfleet "github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/schema/fleet"
)

type FleetHandler struct {
	r *repository.Repository
	v []*pbfleet.Vehicle
}

func NewFleetHandler(r *repository.Repository) (*FleetHandler, error) {
	logger := log.New(os.Stderr, "App info: ", 2)
	data, err := r.GetData()

	if err != nil {
		logger.Fatal(err)
	}

	var vehicles []*pbfleet.Vehicle
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

func (h *FleetHandler) GetVehicles() ([]*pbfleet.Vehicle, error) {
	return h.v, nil
}

// package internal functions
func vehicleFromCsv(s []string) pbfleet.Vehicle {
	logger := log.New(os.Stderr, "App info: ", 2)
	registrationYear, err := strconv.Atoi(s[7])

	if err != nil {
		logger.Fatal(err)
	}

	life, err := strconv.Atoi(s[8])

	if err != nil {
		logger.Fatal(err)
	}

	return pbfleet.Vehicle{
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
