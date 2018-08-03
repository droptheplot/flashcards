package handlers

import "github.com/droptheplot/flashcards/entities"

type DBRepository interface {
	GetSources() ([]entities.Source, error)
	GetSourceByID(ID int) (entities.Source, error)
	GetCardsBySourceID(ID int) ([]entities.Card, error)
	GetUserByEmail(email string) (entities.User, error)
	CreateUser(email string, password string) error
}

type KVRepository interface {
	CreateToken(token string, userID int) error
}

type Handler struct {
	DBRepository DBRepository
	KVRepository KVRepository
}
