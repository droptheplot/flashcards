package handlers

import "github.com/droptheplot/flashcards/repositories/db"

type DBRepository interface {
	GetSources() ([]db.Source, error)
	GetSourceByID(ID int) (db.Source, error)
	GetCardsBySourceID(ID int) ([]db.Card, error)
	GetUserByEmail(email string) (db.User, error)
	CreateUser(email string, password string) error
}

type KVRepository interface {
	CreateToken(token string, userID int) error
}

type Handler struct {
	DBRepository DBRepository
	KVRepository KVRepository
}
