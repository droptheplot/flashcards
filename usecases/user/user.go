package user

import (
	"errors"

	"github.com/droptheplot/flashcards/auth"
	"github.com/droptheplot/flashcards/entities"
)

type DBRepository interface {
	GetUserByEmail(email string) (entities.User, error)
	CreateUser(email string, password string) error
}

type KVRepository interface {
	CreateToken(token string, userID int) error
}

type UseCase struct {
	DBRepository DBRepository
	KVRepository KVRepository
}

func (u *UseCase) CreateUser(email string, password string) error {
	password = auth.HashPassword(password)

	return u.DBRepository.CreateUser(email, password)
}

func (u *UseCase) CreateToken(email string, password string) (string, error) {
	user, err := u.DBRepository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	if user.Password != auth.HashPassword(password) {
		return "", errors.New("password is invalid")
	}

	token, err := auth.GenerateToken()

	if err != nil {
		return "", err
	}

	err = u.KVRepository.CreateToken(token, user.ID)

	return token, err
}
