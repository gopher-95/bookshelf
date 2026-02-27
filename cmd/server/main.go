package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gopher-95/bookshelf/internal/config"
	"github.com/gopher-95/bookshelf/internal/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	err = db.Init(cfg.DBConnectionString())
	if err != nil {
		log.Fatal("Failed to Init database:", err)
	}
	defer db.Close()

	r := chi.NewRouter()
	setupRoutes(r)

	log.Println("Server runnig on port:", cfg.ServerPortString())
	err = http.ListenAndServe(cfg.ServerPortString(), r)
	if err != nil {
		log.Fatal("Failed to run server", err)
	}
}
