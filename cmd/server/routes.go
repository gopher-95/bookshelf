package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func setupRoutes(r *chi.Mux) {
	r.Get("/", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
