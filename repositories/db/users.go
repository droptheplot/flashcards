package db

import (
	"github.com/droptheplot/flashcards/entities"
)

func (r *Repository) CreateUser(email string, password string) error {
	_, err := r.Client.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)

	return err
}

func (r *Repository) GetUserByEmail(email string) (entities.User, error) {
	user := entities.User{}

	err := r.Client.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetUserByID(ID int) (entities.User, error) {
	user := entities.User{}

	err := r.Client.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", ID)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetUserExists(ID int) bool {
	var res bool

	err := r.Client.Get(&res, "SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)", ID)

	if err != nil {
		return false
	}

	return res
}
