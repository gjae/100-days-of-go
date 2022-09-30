package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", (i+1))
	}
}

func main() {
	// usual way for call function
	f("Usual way")

	// Using go rutines
	go f("Using go rutines")

	go func(from string) {
		for i:=0; i < 3; i++ {
			fmt.Println(from, ": ", (i+1))
		}
	}("Go rutines with annonymous func")

	fmt.Println("Sleep")
	time.Sleep(time.Second)
}