package main

import (
	"chat/internal/delivery"
	"chat/internal/repo"
	"chat/internal/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Replace the uri string with your MongoDB deployment's connection string.
	repom := repo.NewPGRepo()
	repom.InitDB()
	defer repom.Close()
	usecase := usecase.NewChatUseCase(repom)
	handler := delivery.NewChatHandler(usecase)

	r := mux.NewRouter()

	r.HandleFunc("/chat/chatlist", handler.GetChatList)
	r.HandleFunc("/chat/conv", handler.GetChat).Methods("GET")
	r.HandleFunc("/chat/put/message", handler.PostMessage)

	srv := &http.Server{Handler: r, Addr: ":8000"}

	log.Fatal(srv.ListenAndServe())

}
