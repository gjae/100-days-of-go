package main

import "fmt"

func rutine1(c chan int) {
	fmt.Println("Rutina 1")
	c <- 5	
}

func rutine2(c chan int) {
	fmt.Println("Rutina 2")
	i, ok := <-c
	if ok {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go rutine2(ch)
	go rutine1(ch)
	rutine1(ch)
}