/**
* The channels can be used for 
* synchronize gorutines
* https://gobyexample.com/channel-synchronization
*/
package main

import (
	"fmt"
	"time"
)

// Worker sleep by 2 seconds and pass
// the true value to done channel
func worker(done chan bool) {
	fmt.Println("Wating...")
	time.Sleep(2)

	fmt.Println("Worker is done")
	done <- true
}


func main() {
	// creating a done channel as bool
	done := make(chan bool)
	go worker(done)

	// This gorutine wait for worker done
	<- done
}