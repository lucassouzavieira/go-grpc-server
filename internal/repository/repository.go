package repository

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