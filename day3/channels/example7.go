package main

import (
	"fmt"
	"sync"
	"time"
)


func worker(id int) {
	fmt.Println("Start Worker ", id)
	time.Sleep(time.Second)

	fmt.Println("Worker ", id, " is done!")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}