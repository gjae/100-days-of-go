package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func echo(w http.ResponseWriter, r *http.Request) {
	con, _ := upgrader.Upgrade(w, r, nil)

	defer con.Close()

	for {
		mt, message, err := con.ReadMessage()

		if err != nil {
			log.Fatal(err)

			break
		}

		fmt.Printf("New message: %s\n", message)

		err = con.WriteMessage(mt, message)
		
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main(){
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8081", nil)
}