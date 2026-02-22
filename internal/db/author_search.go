package db

import "fmt"

func AuthorSearch(author string) ([]Book, error) {
	var book Book

	query := "SELECT * FROM books WHERE author = $1"

	rows, err := db.Query(query, author)
	if err != nil {
		return nil, fmt.Errorf("rows cant be find: %w", err)
	}
	var books []Book

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Pages)
		if err != nil {
			return nil, fmt.Errorf("scan doesnt working: %w", err)
		}
		books = append(books, book)
	}

	return books, nil
}
