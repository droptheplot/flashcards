package source

import "github.com/droptheplot/flashcards/entities"

type DBRepository interface {
	GetSources() ([]entities.Source, error)
	GetSourceByID(ID int) (entities.Source, error)
	GetCardsBySourceID(ID int) ([]entities.Card, error)
}

type UseCase struct {
	DBRepository DBRepository
}

func (u *UseCase) GetSources() ([]entities.Source, error) {
	return u.DBRepository.GetSources()
}

func (u *UseCase) GetSourceByID(ID int) (entities.Source, error) {
	source, err := u.DBRepository.GetSourceByID(ID)

	if err != nil {
		return source, err
	}

	cards, _ := u.DBRepository.GetCardsBySourceID(ID)

	source.Cards = cards

	return source, nil
}
