package handlers

import (
	"github.com/droptheplot/flashcards/repositories/flashcards"
)

type Repository interface {
	GetSources() ([]flashcards.Source, error)
}

type Handler struct {
	Repository Repository
}
