package flashcards

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	DB *sqlx.DB
}

func (r *Repo) Ping() string {
	var pong string

	row := r.DB.QueryRow("SELECT 'pong';")

	if err := row.Scan(&pong); err != nil {
		log.Fatal(err)
	}

	return pong
}
