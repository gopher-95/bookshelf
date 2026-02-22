package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// в глобальной переменной храним бд
var db *sql.DB

// таблица users
const users string = `
CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,	
	login TEXT NOT NULL UNIQUE,
	password_hash TEXT NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
`

// таблица books
const books string = `
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

	// строка подключения
	connectString := "user=postgres password=12345 dbname=bookshelf_db sslmode=disable"
	db, err = sql.Open("postgres", connectString)
	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	}
	// проверка соединения с бд
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	// выполнение запроса на создание таблицы books
	_, err = db.Exec(books)
	if err != nil {
		return fmt.Errorf("unable to build table books: %w", err)
	}

	// выполнение запроса на создание таблицы users
	_, err = db.Exec(users)
	if err != nil {
		return fmt.Errorf("unable to build table users: %w", err)
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
