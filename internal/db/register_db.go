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
