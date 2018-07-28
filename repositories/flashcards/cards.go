package flashcards

type Card struct {
	ID      int    `db:"id" json:"id"`
	Content string `db:"content" json:"content"`
}

func (r *Repository) GetCardsBySourceID(ID int) ([]Card, error) {
	cards := []Card{}

	err := r.DB.Select(&cards, `
		SELECT cards.id, cards.content FROM cards
		INNER JOIN cards_sources
			ON cards.id = cards_sources.card_id
			AND cards_sources.source_id = $1
		ORDER BY cards.id DESC;
	`, ID)

	if err != nil {
		return cards, err
	}

	return cards, nil
}
