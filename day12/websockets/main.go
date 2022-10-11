package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{}

func reverse(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	for {
		mt, message, _ := ws.ReadMessage()
		log.Printf("Message received: %s ", message)

		n := len(message)

		for i := 0; i< n/2; i++ {
			message[i], message[n-1-i] = message[n-1-i], message[i]
		}


		// Response message
		_ = ws.WriteMessage(mt, message)
		log.Printf("Message sent: %s", message)
	}
}

func main() {
	http.HandleFunc("/reverse", reverse)
	log.Fatal(http.ListenAndServe(":5555", nil))
}