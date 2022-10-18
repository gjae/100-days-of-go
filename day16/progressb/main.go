package main

import (
	"fmt"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func run(channel chan <- int, finish chan <- bool, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond)
		channel <- i
	}

	finish <- true
}

func main() {
	count := 10000
	busyChannel := make(chan int)
	finished := make(chan bool)

	bar := pb.StartNew(count)	

	go run(busyChannel, finished, count)

	for i := 0; i < count; i++ {
		select {
		case <- busyChannel:
			bar.Increment()
		case finish := <- finished:
			if finish {
				bar.Finish()
			}
		}
	}

	fmt.Println("Porgressbar finished")
}