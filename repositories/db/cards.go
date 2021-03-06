package db

import "github.com/droptheplot/flashcards/entities"

func (r *Repository) GetCardsBySourceID(ID int) ([]entities.Card, error) {
	cards := []entities.Card{}

	err := r.Client.Select(&cards, `
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

func (r *Repository) AddCardToUser(cardID, userID int, correct bool) error {
	_, err := r.Client.Exec(
		"INSERT INTO cards_users (card_id, user_id, correct) VALUES ($1, $2, $3)",
		cardID, userID, correct,
	)

	return err
}
