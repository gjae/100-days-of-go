package main

import "fmt"

// The syntax var chan <- type mark channel as 
// only send data
func ping(pings chan <- string, msg string) {
	pings <- msg
	// msg := <- pings raise error 
	// because pings channel is only-send data
}

// The syntax var <- chan type mark the channel as
// only receiver
func pong(pings <- chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}


func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Success message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}