package auth

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

	usecase := usecase.NewAuthUseCase(repom, "brr", []byte("brr"), time.Duration(150000000000))
	handler := delivery.NewAuthHandler(usecase)

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
		Addr:    ":8000",
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
