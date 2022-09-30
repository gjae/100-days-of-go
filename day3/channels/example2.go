package main

import "fmt"

func sum(a []int, c chan int) {
	var total int

	for _, val := range a {
		total += val
	}

	c <- total
}

func main() {
	a := []int{1,2,3,4,5,6,7}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	
	fmt.Println(x, y)
}