package main

import "fmt"

func mayPanic() {
	panic("A problem!")
}

func main() {
	
	defer func() {
		if r:= recover(); r != nil {
			fmt.Println("Recover error: \n", r)
		} else {
			fmt.Println("No panics!")
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}