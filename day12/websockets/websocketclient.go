package main 

import (
	"bufio"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)


func main() {
	// Creating url
	// this parsed to string is like: ws://localhost:5555/reverse
	url := url.URL{Scheme: "ws", Host: ":5555",  Path: "/reverse"}

	// Establishing connection
	c, _, _ := websocket.DefaultDialer.Dial(url.String(), nil)

	// Close connection after finished of execute
	defer c.Close()

	// Receive messsages from server
	go func() {
		for {
			_, message, _ := c.ReadMessage()
			log.Printf("Message : %s", message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		c.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))

		log.Printf("Message sent: %s", scanner.Text())
	}
}