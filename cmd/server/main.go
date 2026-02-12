package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gopher-95/bookshelf/internal/db"
)

func main() {
	log.Println("подключение к бд")
	err := db.Init()
	if err != nil {
		log.Fatalf("не удалось подключиться к бд: %v", err)
	}
	defer db.Close()
	log.Println("server is runnug on port :8080")

	r := chi.NewRouter()

	setupRoutes(r)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("server is not working error")
	}

}
