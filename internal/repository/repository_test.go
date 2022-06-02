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
