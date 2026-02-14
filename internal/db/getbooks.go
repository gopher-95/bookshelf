package db

import "fmt"

// функция получения всех книг
func GetBooks() ([]Book, error) {
	var books []Book

	query := "SELECT id, title, author, genre, pages FROM books ORDER by id"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("select query doesnt work: %w", err)
	}

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Pages)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan doesnt work: %w", err)
		}

		books = append(books, book)
	}

	rows.Close()

	return books, nil
}
