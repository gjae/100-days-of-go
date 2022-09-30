package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
    	c <- x
    	x, y = y, x + y
	}
	close(c)
}

func main() {
	ch := make(chan int, 10)

	// go fibonacci(10, ch)
	go fibonacci(cap(ch), ch)

	for i := range ch {
		fmt.Println(i)
	}
}	