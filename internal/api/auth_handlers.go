package api

import (
	"encoding/json"
	"net/http"

	"github.com/gopher-95/bookshelf/internal/db"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user db.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid JSON format")
		return
	}

	if user.Login == "" || user.Password == "" {
		jsonError(w, http.StatusBadRequest, "login or password incorrect")
		return
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "failed tp process password")
		return
	}

	id, err := db.RegistUser(user.Login, hash)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "failed to regist user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "пользователь добавлен",
		"id":      id,
	})
}
