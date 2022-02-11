package main

import (
	delivery2 "auth/internal/delivery"
	repo2 "auth/internal/repo"
	usecase2 "auth/internal/usecase"
	"auth/utils"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
)

func main() {
	connstr, err := utils.ConnStr()
	if err != nil {
		panic(err)
	}
	log.Println(connstr)
	conn, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		panic(err)
	}
	repom := repo2.NewPGRepo(conn)
	usecase := usecase2.NewAuthUseCase(repom)
	handler := delivery2.NewAuthHandler(usecase)

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	{
		s.HandleFunc("/auth/signup", handler.SignUp).Methods("POST")
		s.HandleFunc("/auth/signin", handler.SignIn).Methods("POST")
	}

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
