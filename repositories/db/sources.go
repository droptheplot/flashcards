package db

type Source struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Cards []Card `db:"cards" json:"cards,omitempty"`
}

func (r *Repository) GetSources() ([]Source, error) {
	sources := []Source{}

	err := r.DB.Select(&sources, "SELECT * FROM sources ORDER BY id DESC;")

	if err != nil {
		return sources, err
	}

	return sources, nil
}

func (r *Repository) GetSourceByID(ID int) (Source, error) {
	source := Source{}

	err := r.DB.Get(&source, "SELECT * FROM sources WHERE id = $1 LIMIT 1", ID)

	if err != nil {
		return source, err
	}

	return source, nil
}
