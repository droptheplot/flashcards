package handlers

import (
	"github.com/droptheplot/flashcards/repositories/flashcards"
)

type Repository interface {
	GetSources() ([]flashcards.Source, error)
	GetSourceByID(ID int) (flashcards.Source, error)
	GetCardsBySourceID(ID int) ([]flashcards.Card, error)
}

type Handler struct {
	Repository Repository
}
