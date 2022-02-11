package delivery

import (
	"auth/internal"
	models2 "auth/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

type AuthHandler struct {
	UseCase internal.UseCase
}

func NewAuthHandler(UCase internal.UseCase) *AuthHandler {
	return &AuthHandler{
		UseCase: UCase,
	}
}

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := models2.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response(w, models2.INTERNAL, nil)
		return
	}
	authedUsr, errCode := a.UseCase.SignUp(&user)
	body, err := json.Marshal(authedUsr)
	if err != nil {
		response(w, models2.INTERNAL, nil)
		return
	}
	response(w, errCode, body)
}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	user := models2.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response(w, models2.INTERNAL, nil)
		return
	}
	log.Println(user)
	authedUsr, errCode := a.UseCase.SignIn(&user)

	body, err := json.Marshal(authedUsr)
	if err != nil {
		response(w, models2.INTERNAL, nil)
		return
	}
	response(w, errCode, body)
}

func response(w http.ResponseWriter, errCode models2.ErrorCode, body []byte) {
	switch errCode {
	case models2.OK:
		w.WriteHeader(http.StatusOK)
	case models2.EXISTS:
		w.WriteHeader(http.StatusConflict)
	case models2.NOT_FOUND:
		w.WriteHeader(http.StatusNotFound)
	case models2.CREATED:
		w.WriteHeader(http.StatusCreated)
	case models2.INTERNAL:
		w.WriteHeader(http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(body)
}
