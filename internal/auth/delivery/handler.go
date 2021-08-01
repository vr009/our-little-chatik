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

func MiddleWare(w http.ResponseWriter, r *http.Request) models.User {

	// temporary thing< that is bad
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	var user models.User
	body := make([]byte, 0, 25)

	if _, err := r.Body.Read(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}

	return user
}

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	user := MiddleWare(w, r)

	if mytoken, err := a.UseCase.SignUp(user.UserName, user.Password); err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Print(err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mytoken))
		//r.Method = "GET"
		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	log.Print(user.UserName)

}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {

	user := MiddleWare(w, r)

	if mytoken, err := a.UseCase.SignIn(user.UserName, user.Password); err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Print(err)
	} else {
		//w.WriteHeader(http.StatusOK)
		w.Write([]byte(mytoken))
		r.Method = "GET"
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
