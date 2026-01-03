package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)s

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("WebSocket server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
