package main

import "fmt"

func main() {
	numbers := make(chan int, 2)
	counter := 10

	for i :=0; i < counter; i++ {
		select {
			// This is where the processing take place
		case numbers <- i * i:
			fmt.Println("About to process", i)
		default:
			fmt.Println("No space for ", i, " ")
		}
	}

	fmt.Println()

	for {
		select {
		case num := <-numbers:
			fmt.Println("*", num, "*")
		default:
			fmt.Println("Nothing left to read")
			return
		}
	}
}