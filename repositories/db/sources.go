package db

import "github.com/droptheplot/flashcards/entities"

func (r *Repository) GetSources() ([]entities.Source, error) {
	sources := []entities.Source{}

	err := r.DB.Select(&sources, "SELECT * FROM sources ORDER BY id DESC;")

	if err != nil {
		return sources, err
	}

	return sources, nil
}

func (r *Repository) GetSourceByID(ID int) (entities.Source, error) {
	source := entities.Source{}

	err := r.DB.Get(&source, "SELECT * FROM sources WHERE id = $1 LIMIT 1", ID)

	if err != nil {
		return source, err
	}

	return source, nil
}
