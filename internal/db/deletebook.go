package db

import "fmt"

func DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = $1"

	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete query doesnt work: %w", err)
	}

	return nil
}
