package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"our-little-chatik/internal/auth/delivery"
	"our-little-chatik/internal/auth/repo"
	"our-little-chatik/internal/auth/usecase"
	"time"
)

func main() {
	//	repom := repo.NewmockRepo()
	repom := repo.NewPGRepo()
	repom.InitDB()
	defer repom.Close()

	usecase := usecase.NewAuthUseCase(repom, "brr", []byte("brr"), time.Duration(15))
	handler := delivery.NewAuthHandler(usecase)

	r := mux.NewRouter()
	r.HandleFunc("/auth/signup", handler.SignUp).Methods("POST")
	r.HandleFunc("/auth/signin", handler.SignIn).Methods("POST")
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Got it"))
	}).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
