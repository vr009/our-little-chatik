package delivery

import (
	"encoding/json"
	"fmt"
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

func MiddleWare(w http.ResponseWriter, r *http.Request) models.User {
	// temporary thing
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	var user models.User
	body := make([]byte, 0, 25)

	if _, err := r.Body.Read(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil || user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}

	return user
}

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	user := MiddleWare(w, r)

	token, err := a.UseCase.SignUp(user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Print(err)
	} else {
		w.Header().Set("Set-Cookie", fmt.Sprintf("ssid=%s; path=/; HttpOnly", token))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	log.Print(user.Username)

}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {

	user := MiddleWare(w, r)

	if token, err := a.UseCase.SignIn(user); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Print(err)
	} else {
		w.Header().Set("Set-Cookie", fmt.Sprintf("ssid=%s; path=/; HttpOnly", token))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
