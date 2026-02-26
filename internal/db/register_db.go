package db

import (
	"fmt"
)

func RegistUser(login, hashPassword string) (int, error) {
	var id int

	query := "INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id"

	err := db.QueryRow(query, login, hashPassword).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("regist user error: %w", err)
	}

	return id, nil
}

func CheckLogin(login string) (bool, error) {
	var foundLogin bool

	query := "SELECT EXISTS (SELECT 1 FROM users WHERE login = $1)"

	err := db.QueryRow(query, login).Scan(&foundLogin)
	if err != nil {
		return false, fmt.Errorf("database error: %w", err)
	}

	return foundLogin, nil
}

func GetPasswordHash(login string) ([]byte, error) {
	var password string
	query := "SELECT password_hash FROM users WHERE login = $1"

	err := db.QueryRow(query, login).Scan(&password)
	if err != nil {
		return nil, fmt.Errorf("get password error: %w", err)
	}

	passwordInBytes := []byte(password)

	return passwordInBytes, nil
}
