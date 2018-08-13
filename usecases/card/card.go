package card

type DBRepository interface {
	AddCardToUser(cardID, userID int, correct bool) error
}

type UseCase struct {
	DBRepository DBRepository
}

func (u *UseCase) AddCardToUser(cardID, userID int, correct bool) error {
	return u.DBRepository.AddCardToUser(cardID, userID, correct)
}
