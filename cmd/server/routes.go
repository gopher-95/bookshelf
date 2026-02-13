package main

import (
	"github.com/go-chi/chi"
	"github.com/gopher-95/bookshelf/internal/api"
)

func setupRoutes(r *chi.Mux) {
	r.Post("/", api.AddBookHandler)
}
