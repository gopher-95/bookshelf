package main

import (
	"github.com/go-chi/chi"
	"github.com/gopher-95/bookshelf/internal/api"
)

func setupRoutes(r *chi.Mux) {
	r.Post("/add", api.AddBookHandler)
	r.Post("/registration", api.RegisterHandler)
	r.Get("/books", api.GetAllBooks)
	r.Get("/books/{id}", api.GetBook)
	r.Get("/books/search", api.SearchBook)
	r.Delete("/books/{id}", api.DeleteBook)
}
