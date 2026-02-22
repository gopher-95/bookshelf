package main

import (
	"github.com/go-chi/chi"
	"github.com/gopher-95/bookshelf/internal/api"
)

func setupRoutes(r *chi.Mux) {
	r.Post("/add", api.AddBookHandler)
	r.Get("/books", api.GetAllBooks)
	r.Get("/books/{id}", api.GetBook)
	r.Delete("/books/{id}", api.DeleteBook)
	r.Get("/books/search", api.SearchBook)
	r.Post("/registration", api.RegisterHandler)
}
