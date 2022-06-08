package service

import (
	"bytes"
	"testing"

	r "github.com/lucassouzavieira/go-grpc-server/internal/repository"
)

var (
	mockup = "../../test/data/vehicles.csv"
)

func TestGetVehicles(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("any,csv,data")

	repository := r.NewRepository(mockup)

	handler, err := NewFleetHandler(repository)

	if err != nil {
		t.Error("Failed to create FleetHandler")
	}

	vehicles, err := handler.GetVehicles()

	if err != nil {
		t.Error("Failed to load data")
	}

	var expected int32 = 3

	if len(vehicles) != 3 {
		t.Errorf("Expected %d lines, found %d", expected, len(vehicles))
	}
}
