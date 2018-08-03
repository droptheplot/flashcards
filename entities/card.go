package entities

type Card struct {
	ID      int    `db:"id" json:"id"`
	Content string `db:"content" json:"content"`
}
