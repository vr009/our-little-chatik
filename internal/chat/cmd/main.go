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
	r.HandleFunc("/fetch/chatlist", handler.GetChatList)
	r.HandleFunc("/fetch/conv", handler.GetChat)
	r.HandleFunc("/put/message", handler.PostMessage)

	srv := &http.Server{Handler: r, Addr: "8080"}

	log.Fatal(srv.ListenAndServe())

}
