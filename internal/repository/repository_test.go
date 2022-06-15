package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mockup       = "../../test/data/mockup.csv"
	mockup_write = "../../test/data/mockup_write.csv"
)

func TestReadFile(t *testing.T) {
	repository := NewRepository(mockup)

	_, err := repository.GetData()
	if err != nil {
		t.Error("Failed to read csv")
	}
}

func TestGetHeaders(t *testing.T) {
	repository := NewRepository(mockup)

	headers, err := repository.GetHeaders()
	if err != nil {
		t.Error("Failed to read csv")
	}

	assert.Equal(t, "header1", headers[0])
}

func TestWriteFile(t *testing.T) {
	repository := NewRepository(mockup_write)

	line := []string{"yet", "another", "line"}
	err := repository.AddData(line)

	if err != nil {
		t.Error("Failed to read csv")
	}
}

func TestSingleFilter(t *testing.T) {
	repository := NewRepository(mockup)

	var filterA = Filter{
		property: "header3",
		value:    "data",
	}

	var filters = make([]*Filter, 0)
	filters = append(filters, &filterA)

	filtered, err := repository.Filter(filters)

	if err != nil {
		t.Error("Failed to filter data")
	}

	assert.Equal(t, 3, len(filtered))
	assert.Equal(t, "data", filtered[0][2])
}

func TestMultipleFilters(t *testing.T) {
	repository := NewRepository(mockup)

	var filterA = Filter{
		property: "header3",
		value:    "ssj3",
	}

	var filterB = Filter{
		property: "header2",
		value:    "fake",
	}

	var filters = make([]*Filter, 0)
	filters = append(filters, &filterA)
	filters = append(filters, &filterB)

	filtered, err := repository.Filter(filters)

	if err != nil {
		t.Error("Failed to filter data")
	}

	assert.Equal(t, 1, len(filtered))
}

func TestCount(t *testing.T) {
	repository := NewRepository(mockup)

	count, err := repository.Count()
	if err != nil {
		t.Error("Failed to read csv")
	}

	assert.Equal(t, 7, count)
}
