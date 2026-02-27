package db

import "fmt"

// заносим пользователя в бд при регистрации
func RegisterUser(login, hashPassword string) (int64, error) {
	var id int64

	query := "INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id"

	err := db.QueryRow(query, login, hashPassword).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Could not return id")
	}

	return id, nil
}

// получаем по логину пользователя его id и захэшированный пароль
func LoginUser(login string) (*User, error) {
	user := &User{}

	query := "SELECT id, login, password_hash FROM users WHERE login = $1"

	row := db.QueryRow(query, login)
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("check user: %w", err)
	}

	return user, nil
}
