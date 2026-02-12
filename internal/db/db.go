package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// в глобальной переменной храним бд
var db *sql.DB

// создание таблицы books
const schema string = `
CREATE TABLE IF NOT EXISTS books (
	id BIGSERIAL PRIMARY KEY,
	title TEXT NOT NULL, 
	author TEXT NOT NULL,
	genre TEXT,
	pages INTEGER CHECK (pages > 0)
	);
`

// подключение к бд и создание таблицы
func Init() error {
	var err error
	connectString := "user=postgres password=12345 dbname=bookshelf_db sslmode=disable"
	db, err = sql.Open("postgres", connectString)
	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		return fmt.Errorf("unable to build database")
	}

	return nil
}

// закрытие бд
func Close() error {
	if db != nil {
		db.Close()
	}

	return nil
}
