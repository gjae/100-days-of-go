package main

import "fmt"
import "time"
import "os"
import "strconv"

func main() {
	count := 10
	if len(os.Args) > 1 {
		count, _ = strconv.Atoi(os.Args[1])
	}

	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")

}