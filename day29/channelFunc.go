package main

import "fmt"
import "sync"

var wg sync.WaitGroup

// Arrow before chan keyword = channel for readonly
// Arrow between chan keyword and data type = writeonly channel
func wrChannel(out <-chan int, in chan <- int) {
	defer wg.Done()
	x := <- out
	fmt.Println(x)

	in <- x * x
}

func main() {
	ch := make(chan int)
	in := make(chan int)

	wg.Add(1)
	go wrChannel(ch, in)

	ch <- 3
	fmt.Println(<-in)
	wg.Wait()
	
	close(ch)
	close(in)
}