package main

import (
	"log"
	"net/http"
	"os"

	"github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(socketio.DefaultConfig)
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	port := os.Getenv("PORT")
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Printf("Serving at localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
