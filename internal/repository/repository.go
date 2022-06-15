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

type Filter struct {
	Property string
	Value    string
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

func (r *Repository) GetHeaders() ([]string, error) {
	data, err := r.GetData()

	if err != nil {
		log.WithError(err).Error("Error while trying to read data")
		return nil, err
	}

	headers := make([][]string, 0)

	for i, v := range data {
		headers = append(headers, v)

		if i > 0 {
			break
		}
	}

	return headers[0], nil
}

func (r *Repository) AddData(l []string) error {
	f, err := os.OpenFile(r.filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

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

	if err := csvWriter.Error(); err != nil {
		log.Fatal(err)
	}

	return err
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

func (r *Repository) Filter(fs []*Filter) ([][]string, error) {
	data, err := r.GetData()

	if err != nil {
		log.WithError(err).Error("Error while trying to read data")
	}

	headers, err := r.GetHeaders()
	if err != nil {
		log.WithError(err).Error("Error while trying to read data")
	}

	getFilterIndex := func(headers []string, fs *Filter) int32 {
		var index int32 = 0
		for i, header := range headers {
			if header == fs.Property {
				index = int32(i)
				break
			}
		}

		return index
	}

	for _, f := range fs {
		filtered := make([][]string, 0)
		var filterIndex = getFilterIndex(headers, f)

		// Check the data to filter based on the current filter
		for _, entry := range data {
			if entry[filterIndex] == f.Value {
				filtered = append(filtered, entry)
			}
		}

		data = filtered
	}

	return data, nil
}
