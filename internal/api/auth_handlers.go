package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gopher-95/bookshelf/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var regUser db.LoginRequest

	user, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid login or password")
		return
	}

	err = json.Unmarshal(user, &regUser)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Could not parse request data")
		return
	}

	hasedPassword, err := hashPassword(regUser.Password)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Could not hash password")
		return
	}

	id, err := db.RegisterUser(regUser.Login, hasedPassword)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Counld not get id")
		return
	}

	jsonResponse(w, http.StatusCreated, "Пользователь успешно добавлен", "ID", id)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest db.LoginRequest

	user, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid credentials")
		return
	}

	err = json.Unmarshal(user, &loginRequest)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "could not parse request data")
		return
	}

	loginUser, err := db.LoginUser(loginRequest.Login)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.PasswordHash), []byte(loginRequest.Password))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid credentials")
		return
	}

	jsonResponse(w, http.StatusOK, "Вы успешно вошли!", "id", loginUser.ID)
}
