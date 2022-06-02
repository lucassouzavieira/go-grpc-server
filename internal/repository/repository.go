package repository

import (
	"encoding/csv"
	"log"
	"os"
)

// Loads and handle data from our dataset

type Repository struct {
	filepath string
	filetype string
}

func NewRepository(filepath string) *Repository {
	return &Repository{
		filepath: filepath,
		filetype: "csv",
	}
}

func (r *Repository) Filepath() string {
	return r.filepath
}

func (r *Repository) Type() string {
	return r.filetype
}

func (r *Repository) GetData() ([][]string, error) {
	logger := log.New(os.Stderr, "repository: ", 2)
	f, err := os.Open(r.filepath)
	if err != nil {
		logger.Fatalln(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	return data, err
}