package db

import "fmt"

func GetBook(id int) (Book, error) {
	var book Book

	query := "SELECT * FROM	 books WHERE id = $1"

	row := db.QueryRow(query, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Pages)
	if err != nil {
		return Book{}, fmt.Errorf("get query doesnt work: %w", err)
	}

	return book, nil
}
