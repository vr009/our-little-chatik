package main

import (
	_ "auth/docs"
	delivery2 "auth/internal/delivery"
	repo2 "auth/internal/repo"
	usecase2 "auth/internal/usecase"
	"auth/middleware"
	"auth/utils"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title           Auth API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  slavarianov@yandex.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
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
	encr := usecase2.Encrypter_pbkdf2{}
	repom := repo2.NewPGRepo(conn, encr)
	usecase := usecase2.NewAuthUseCase(repom)
	handler := delivery2.NewAuthHandler(usecase)

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	{
		s.HandleFunc("/auth/signup", handler.SignUp).Methods("POST")
		s.HandleFunc("/auth/signin", handler.SignIn).Methods("POST")
	}
	r.Use(middleware.CORSMiddleware)
	r.PathPrefix("/documentation").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
