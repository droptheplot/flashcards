package db

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Client *sqlx.DB
}
