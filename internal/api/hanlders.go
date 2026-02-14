package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gopher-95/bookshelf/internal/db"
)

// фукнция получения книги из бд
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := db.GetBooks()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "cant get data from database")
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "книги успешно получены",
		"books":   books,
	})
}

// функция добавления книги в бд
func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var book db.Book
	data, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		jsonError(w, http.StatusInternalServerError, "unable to read the request body")
		return
	}

	err = json.Unmarshal(data, &book)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid JSON format")
		return
	}

	if book.Title == "" || book.Author == "" {
		jsonError(w, http.StatusBadRequest, "title and author cant be empty")
		return
	}

	err = db.AddBook(&book)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "failed to add book")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "книга успешно добавлена",
		"book":    book,
	})

}

// функция для формирования json ответа при ошибках
func jsonError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
