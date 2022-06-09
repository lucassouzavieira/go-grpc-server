package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	r "github.com/lucassouzavieira/go-grpc-server/internal/repository"
)

var (
	incidents_mockup = "../../test/data/incidents.csv"
)

func TestGetIncidents(t *testing.T) {
	repository := r.NewRepository(incidents_mockup)
	handler, err := NewIncidentHandler(repository)

	if err != nil {
		t.Error("Failed to create IncidentHandler")
	}

	incidents, err := handler.GetIncidents()

	if err != nil {
		t.Error("Failed to load data")
	}

	assert.Equal(t, 4, len(incidents))
}

func TestGetIncidentsByAnimalGroup(t *testing.T) {
	repository := r.NewRepository(incidents_mockup)
	handler, err := NewIncidentHandler(repository)

	if err != nil {
		t.Error("Failed to create IncidentHandler")
	}

	catIncidents, err := handler.GetIncidentsByAnimalGroup("Cat")

	if err != nil {
		t.Error("Failed to load data")
	}

	assert.Equal(t, 0, len(catIncidents))

	dogIncidents, err := handler.GetIncidentsByAnimalGroup("Dog")

	if err != nil {
		t.Error("Failed to load data")
	}

	assert.Equal(t, 2, len(dogIncidents))
}
