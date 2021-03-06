package service

// Loads data from a csv entry and transform to a proto model struct
import (
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"

	repository "github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
)

type FleetHandler struct {
	r *repository.Repository
	v []*fleet.Vehicle
}

type FleetStats struct {
	Active     int32
	Reserve    int32
	Training   int32
	AverageAge float32
}

func NewFleetHandler(r *repository.Repository) (*FleetHandler, error) {
	data, err := r.GetData()

	if err != nil {
		log.Fatal(err)
		return nil, err
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

func (h *FleetHandler) Filter(fs []*repository.Filter) (bool, error) {
	data, err := h.r.Filter(fs)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	var vehicles []*fleet.Vehicle
	for i, v := range data {
		if i == 0 {
			continue // skip headers
		}

		vehicle := vehicleFromCsv(v)
		vehicles = append(vehicles, &vehicle)
	}

	h.v = vehicles
	return true, nil
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

func (h *FleetHandler) GetVehiclesByYear(year int32) ([]*fleet.Vehicle, error) {
	var filtered []*fleet.Vehicle = make([]*fleet.Vehicle, 0)

	for _, v := range h.v {
		if v.RegistrationYear == year {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}

func (h *FleetHandler) GetStats(fs []*repository.Filter) (*FleetStats, error) {
	_, err := h.Filter(fs)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var stats = FleetStats{}
	var agesSum int32 = 0

	for _, v := range h.v {
		if v.GetOperationalStatus() == "ACTIVE" {
			stats.Active = stats.Active + 1
		}

		if v.GetOperationalStatus() == "TRAINING" {
			stats.Training = stats.Training + 1
		}

		if v.GetOperationalStatus() == "RESERVE" {
			stats.Reserve = stats.Reserve + 1
		}

		vxAge := int32(time.Now().Year()) - v.GetRegistrationYear()
		agesSum += vxAge
	}

	stats.AverageAge = 0
	var denom int32 = int32(len(h.v))

	if denom > 0 {
		stats.AverageAge = float32(agesSum / denom)
	}

	for _, f := range fs {
		log.Infof("Filters: %+v", *f)
	}

	log.Infof("Stats: %+v", stats)
	return &stats, nil
}

func (h *FleetHandler) SaveVehicle(v *fleet.Vehicle) (bool, error) {
	line, err := toCsv(v)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	err = h.r.AddData(line)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

// package internal functions
func toCsv(v *fleet.Vehicle) ([]string, error) {
	representation := make([]string, 9)

	representation[0] = v.GetFleetNumber()
	representation[1] = v.GetOperationalStatus()
	representation[2] = v.GetLfb()
	representation[3] = v.GetMake()
	representation[4] = v.GetModel()
	representation[5] = v.GetType()
	representation[6] = v.GetCategory()
	representation[7] = decimal.NewFromInt32(v.GetRegistrationYear()).StringFixed(0)
	representation[8] = decimal.NewFromInt32(v.GetLife()).StringFixed(0)

	return representation, nil
}

func vehicleFromCsv(s []string) fleet.Vehicle {
	registrationYear, err := decimal.NewFromString(s[7])

	if err != nil {
		log.Fatal(err)
	}

	life, err := decimal.NewFromString(s[8])

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
		RegistrationYear:  int32(registrationYear.IntPart()),
		Life:              int32(life.IntPart()),
	}
}
