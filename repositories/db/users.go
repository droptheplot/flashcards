package db

import "github.com/droptheplot/flashcards/entities"

func (r *Repository) CreateUser(email string, password string) error {
	_, err := r.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)

	return err
}

func (r *Repository) GetUserByEmail(email string) (entities.User, error) {
	user := entities.User{}

	err := r.DB.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)

	if err != nil {
		return user, err
	}

	return user, nil
}
