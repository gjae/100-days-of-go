package main

import "fmt"

func main() {
	ch := make(chan int, 2) // if turn 2 by 1 raise error, but with 3 not 

	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}