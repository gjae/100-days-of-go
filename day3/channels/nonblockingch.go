package main

import "fmt"

func main() {
	messages := make(chan string, 1)

	select {
	case message := <- messages :
		fmt.Printf("Message received: %v", message)

	default:
		fmt.Println("No message receiver")
	}


	msg := "Send"

	select {
	case messages <- msg:
		fmt.Printf("SEnd: %v", <-messages)
	default:
		fmt.Println("No received")
	}
}