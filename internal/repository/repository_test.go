package repository

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	mockup = "../../test/data/mockup.csv"
)

func TestReadFile(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("any,csv,data")

	repository := NewRepository(mockup)

	content, err := repository.GetData()
	if err != nil {
		t.Error("Failed to read csv")
	}

	fmt.Print(content)
}

func TestCount(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("any,csv,data")

	repository := NewRepository(mockup)

	count, err := repository.Count()
	if err != nil {
		t.Error("Failed to read csv")
	}

	var expected int32 = 3

	if count != 3 {
		t.Errorf("Expected %d lines, found %d", expected, count)
	}
}
