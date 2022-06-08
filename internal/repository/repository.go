package repository

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
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
	f, err := os.Open(r.filepath)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	return data, err
}

func (r *Repository) AddData(l []string) error {
	f, err := os.Open(r.filepath)

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	csvWriter := csv.NewWriter(f)
	err = csvWriter.Write(l)

	if err != nil {
		log.Fatalln(err)
	}

	csvWriter.Flush()
	return nil
}

func (r *Repository) Count() (int, error) {
	f, err := os.Open(r.filepath)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	return len(data), err
}
