package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gopher-95/bookshelf/internal/db"
	"golang.org/x/crypto/bcrypt"
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
	jsonResponse(w, http.StatusOK, "пользователь добавлен", "id", id)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest db.LoginRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid body")
		return
	}

	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "")
	}

	correctLogin, err := db.CheckLogin(loginRequest.Login)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if correctLogin {
		hashedPassword, err := db.GetPasswordHash(loginRequest.Login)
		if err != nil {
			jsonError(w, http.StatusInternalServerError, "database error")
			return
		}
		err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(loginRequest.Password))
		if err != nil {
			jsonError(w, http.StatusBadRequest, "incorrect password")
			return
		}
	}

}
