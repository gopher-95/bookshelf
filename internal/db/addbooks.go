package db

import "fmt"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Pages  int    `json:"pages"`
}

// добавление книги в бд и получение индентификатора последней добавленной книги
func AddBook(book *Book) error {
	query := "INSERT INTO books (title, author, genre, pages) VALUES ($1, $2, $3, $4) RETURNING id"

	err := db.QueryRow(query, book.Title, book.Author, book.Genre, book.Pages).Scan(&book.ID)
	if err != nil {
		return fmt.Errorf("add book: %w", err)
	}

	return nil

}
