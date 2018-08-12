package user

import (
	"errors"

	"github.com/droptheplot/flashcards/auth"
	"github.com/droptheplot/flashcards/entities"
)

type DBRepository interface {
	GetUserByEmail(email string) (entities.User, error)
	GetUserExists(userID int) bool
	CreateUser(email string, password string) error
}

type KVRepository interface {
	CreateToken(token string, userID int) error
}

type UseCase struct {
	DBRepository DBRepository
	KVRepository KVRepository
	Secret       string
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

	token, err := auth.GenerateToken(user.ID, u.Secret)

	if err != nil {
		return "", err
	}

	err = u.KVRepository.CreateToken(token, user.ID)

	return token, err
}

func (u *UseCase) AuthenticateUser(token string) (int, error) {
	userID, err := auth.ParseToken(token, u.Secret)

	if err != nil {
		return 0, err
	}

	if !u.DBRepository.GetUserExists(userID) {
		return 0, errors.New("user doesn't exist")
	}

	return userID, nil
}
