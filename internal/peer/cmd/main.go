package main

import (
	"flag"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"our-little-chatik/internal/peer/internal/delivery"
	repo2 "our-little-chatik/internal/peer/internal/repo"
	usecase2 "our-little-chatik/internal/peer/internal/usecase"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

var defaultServer = "127.0.0.1:3301"
var defaultOpts = tarantool.Opts{
	Timeout: 500 * time.Millisecond,
	User:    "test",
	Pass:    "test",
	//Concurrency: 32,
	//RateLimit: 4*1024,
}

func main() {
	conn, err := tarantool.Connect(defaultServer, defaultOpts)
	if err != nil {
		panic("failed to connect to tarantool")
	}
	defer conn.Close()
	repo := repo2.NewTarantoolRepo(conn)
	messageManager := usecase2.NewMessageManager(repo)
	usecase := usecase2.NewPeerUsecaseImpl()
	peerServer := delivery.NewPeerServer(usecase, messageManager)

	go messageManager.Work()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		peerServer.WSServe(w, r)
	})
	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
