package main

import (
	"chat_app_backend/pkg/websocket"
	"fmt"
	"net/http"
)

func serverWS(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writter(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serverWS)
}

func main() {
	fmt.Println("Distributed Chat App Versi 1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
