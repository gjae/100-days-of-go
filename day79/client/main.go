package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

const SOCKET_PATH = "/tmp/testsocket.sock"

func main() {
	conn, err := net.Dial("unix", SOCKET_PATH)
	defer conn.Close()

	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}

	message := []byte("Hello world")

	conn.Write(message)

	for {
		var message string
		fmt.Print("Mensaje: ")
		fmt.Scan(&message)
		fmt.Println("")
		conn.Write([]byte(message))
		if strings.ToUpper(message) == "EXIT" {
			break
		}
	}
}
