package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"our-little-chatik/internal/auth"
	"our-little-chatik/internal/models"
)

type AuthHandler struct {
	UseCase auth.UseCase
}

func NewAuthHandler(UCase auth.UseCase) *AuthHandler {
	return &AuthHandler{
		UseCase: UCase,
	}
}

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	var user models.User
	body := make([]byte, 0, 25)

	if _, err := r.Body.Read(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}

	if err := a.UseCase.SignUp(user); err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Print(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r http.Request) {
	var user models.User
	body := make([]byte, 0, 25)

	if _, err := r.Body.Read(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}

	if err := a.UseCase.SignIn(user); err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Print(err)
	}

	w.WriteHeader(http.StatusOK)
}
