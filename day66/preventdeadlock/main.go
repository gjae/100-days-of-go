package main

import (
	"fmt"
	"runtime"
	"time"
)

var queue = make(chan string)

func worker(ch chan string, worker int) {
	for v := range ch {
		fmt.Printf("Worker %d: %s\n", worker, v)

	}
}

func Enqueue(ch chan string) {
	defer close(ch)
	defer close(queue)

	// Loop is brokon when pass 3 seconds without receive
	// news items in queue
	for {
		select {
		case <-time.After(time.Minute):
			break
		case d := <-queue:
			ch <- d
		}

	}
}

func main() {
	runtime.GOMAXPROCS(3)
	ch := make(chan string, 2)
	inputs := []string{
		"First",
		"Second",
		"Third",
		"Fourth",
		"Fifth",
		"Sixth",
	}

	go Enqueue(ch)

	// Launch two workers
	for i := 0; i < 2; i++ {
		go worker(ch, i)
	}

	// Read from data source and enqueue
	// for the workers take jobs
	for _, input := range inputs {
		queue <- input
		time.Sleep(time.Second)
	}

}
