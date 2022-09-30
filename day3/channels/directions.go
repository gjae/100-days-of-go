/*
* learn take from https://gobyexample.com/channel-directions
*/
package main

import "fmt"

// Receive channel only send (write-only)
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// Receive channel for read and write
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}