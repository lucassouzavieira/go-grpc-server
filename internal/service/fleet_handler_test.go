package service

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/stretchr/testify/assert"

	r "github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/fleet"
)

var (
	mockup = "../../test/data/vehicles.csv"
)

func v(status, make, category string) (*fleet.Vehicle, error) {
	return &fleet.Vehicle{
		FleetNumber:       rand.String(5),
		OperationalStatus: status,
		Lfb:               "LFB",
		Make:              make,
		Model:             rand.String(8),
		Type:              rand.String(8),
		Category:          category,
		RegistrationYear:  2014,
		Life:              10,
	}, nil
}

func TestGetVehicles(t *testing.T) {
	repository := r.NewRepository(mockup)
	handler, err := NewFleetHandler(repository)

	if err != nil {
		t.Error("Failed to create FleetHandler")
	}

	vehicles, err := handler.GetVehicles()

	if err != nil {
		t.Error("Failed to load data")
	}

	assert.Equal(t, 3, len(vehicles))
}

func TestSaveVehicle(t *testing.T) {
	repository := r.NewRepository(mockup)
	handler, err := NewFleetHandler(repository)

	if err != nil {
		t.Error("Failed to create FleetHandler")
	}

	vehicles, err := handler.GetVehicles()

	if err != nil {
		t.Error("Failed to load data")
	}

	assert.Equal(t, 3, len(vehicles))

	vxCase, err := v("ACTIVE", "Mercedes", "N3")
	if err != nil {
		t.Error("Failed to create vehicle data")
	}

	reserve, err := handler.SaveVehicle(vxCase)

	if err != nil {
		t.Error("Failed to save data")
	}

	assert.Equal(t, true, reserve)
}

func TestGetVehiclesByOperationalStatus(t *testing.T) {
	repository := r.NewRepository(mockup)
	handler, err := NewFleetHandler(repository)

	if err != nil {
		t.Error("Failed to create FleetHandler")
	}

	active, err := handler.GetVehiclesByOperationalStatus("TRAINING")
	if err != nil {
		t.Error("Failed to load data")
	}

	assert.Equal(t, 1, len(active), fmt.Sprintf("%+v", active))

	reserve, err := handler.GetVehiclesByOperationalStatus("RESERVE")
	if err != nil {
		t.Error("Failed to load data")
	}

	assert.Equal(t, 1, len(reserve), fmt.Sprintf("%+v", reserve))
}
