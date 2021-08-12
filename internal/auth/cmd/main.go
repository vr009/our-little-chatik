package main

import (
	delivery2 "auth/internal/delivery"
	repo2 "auth/internal/repo"
	usecase2 "auth/internal/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	//	repom := repo.NewmockRepo()
	repom := repo2.NewPGRepo()
	repom.InitDB()
	defer repom.Close()

	usecase := usecase2.NewAuthUseCase(repom, "brr", []byte("brr"), time.Duration(150000000000))
	handler := delivery2.NewAuthHandler(usecase)

	r := mux.NewRouter()
	r.HandleFunc("/auth/signup", handler.SignUp).Methods("POST")
	r.HandleFunc("/auth/signin", handler.SignIn).Methods("POST")

	s := r.PathPrefix("").Subrouter()
	s.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Got it"))
	}).Methods("GET")
	s.Use(usecase.AuthMiddleWare)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
