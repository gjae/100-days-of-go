package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	defer close(ch)

	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("Successfully wrote", i, " to ch")
	}

}

func main() {
	// creates capacity of 2

	ch := make(chan int, 2)
	go worker(ch)

	time.Sleep(time.Second * 2)

	for v := range ch {
		fmt.Println("Read value ", v, " from ch")
		time.Sleep(time.Second * 2)
	}
}
