package db

import (
	"database/sql"
	"fmt"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Pages  int    `json:"pages"`
}

// добавление книги в бд и получение индентификатора последней добавленной книги
func AddBook(book *Book) (int64, error) {
	query := "INSERT INTO books (title, author, genre, pages) VALUES (:title, :author, :genre, :pages)"

	result, err := db.Exec(query,
		sql.Named("title", book.Title),
		sql.Named("author", book.Title),
		sql.Named("genre", book.Title),
		sql.Named("pages", book.Title))
	if err != nil {
		return 0, fmt.Errorf("add book: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last id: %w", err)
	}

	return id, nil
}
