package repository

import (
	"fmt"
	"testing"
)

var (
	mockup = "../../test/data/mockup.csv"
	mockup_write = "../../test/data/mockup_write.csv"
)

func TestReadFile(t *testing.T) {
	repository := NewRepository(mockup)

	content, err := repository.GetData()
	if err != nil {
		t.Error("Failed to read csv")
	}

	fmt.Print(content)
}

func TestWriteFile(t *testing.T) {
	repository := NewRepository(mockup_write)

	line := []string{"yet", "another", "line"}
	err := repository.AddData(line)

	if err != nil {
		t.Error("Failed to read csv")
	}
}

func TestCount(t *testing.T) {
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
