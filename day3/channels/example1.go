package main

import "fmt"

func main() {
	// Message var is created as a channel of string type
	message := make(chan string)

	// A gorutine is runned and pass "ping" string
	// to message channel
	go func() { message <- "ping" }()

	// <- channel sintax retrieve a value that was send from 
	// gorutine 
	msg := <-message
	fmt.Println(msg)
}