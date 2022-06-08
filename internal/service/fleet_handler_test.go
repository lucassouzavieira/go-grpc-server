package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	r "github.com/lucassouzavieira/go-grpc-server/internal/repository"
)

var (
	mockup = "../../test/data/vehicles.csv"
)

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

func TestGetVehiclesByOperationalStatus(t *testing.T) {
	repository := r.NewRepository(mockup)
	handler, err := NewFleetHandler(repository)

	if err != nil {
		t.Error("Failed to create FleetHandler")
	}

	active, err := handler.GetVehiclesByOperationalStatus("ACTIVE")
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
