package delivery

import (
	"auth/internal"
	models2 "auth/internal/models"
	"encoding/json"
	"fmt"
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

func MiddleWare(w http.ResponseWriter, r *http.Request) models2.User {
	// temporary thing
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	var user models2.User
	body := make([]byte, 0, 25)

	log.Print(r.RequestURI)

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

func (a *AuthHandler) GetUsersList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	List, err := a.UseCase.FetchUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(List)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resp)
	w.WriteHeader(http.StatusOK)

}
