package main

import "fmt"
import "time"

func main() {
	go func(x int) {
		fmt.Printf("%d ", x)
	}(10)

	time.Sleep(time.Second)

	fmt.Println("Exiting...")
}