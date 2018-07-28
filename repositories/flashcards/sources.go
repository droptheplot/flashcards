package flashcards

type Source struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

func (r *Repository) GetSources() ([]Source, error) {
	sources := []Source{}

	err := r.DB.Select(&sources, "SELECT * FROM sources ORDER BY id DESC;")

	if err != nil {
		return sources, err
	}

	return sources, nil
}
