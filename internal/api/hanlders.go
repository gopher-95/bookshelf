package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gopher-95/bookshelf/internal/db"
)

func SearchBook(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author == "" {
		jsonError(w, http.StatusBadRequest, "invalid name of author")
		return
	}

	books, err := db.AuthorSearch(author)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "failed to get books")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "книги по автору успешно получены",
		"books":   books,
	})
}

// функция получения одной книги
func GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		jsonError(w, http.StatusBadRequest, "missing Book ID")
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid book ID")
		return
	}

	book, err := db.GetBook(intID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "failed to get book")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "книга успешно получена",
		"book":    book,
	})
}

// функция удаления книги из бд
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		jsonError(w, http.StatusBadRequest, "missing Book ID")
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid book ID")
		return
	}

	err = db.DeleteBook(intID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "failed to delete book")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "книга успешно удалена",
		"id":      intID,
	})

}

// фукнция получения книг из бд
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
		"book":    book.Title,
	})
}
