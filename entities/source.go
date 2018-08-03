package entities

type Source struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Cards []Card `db:"cards" json:"cards,omitempty"`
}
