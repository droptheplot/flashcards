package db

type User struct {
	ID       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}

func (r *Repository) CreateUser(email string, password string) error {
	_, err := r.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)

	return err
}

func (r *Repository) GetUserByEmail(email string) (User, error) {
	user := User{}

	err := r.DB.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)

	if err != nil {
		return user, err
	}

	return user, nil
}
