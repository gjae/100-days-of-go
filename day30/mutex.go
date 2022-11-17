package main

import (
	"fmt"
	"os"
	"sync"
	"time"
	"strconv"
)

// Mutex: Mutual Exclusion Variable
// Avoid race condition
var mutex sync.Mutex
var v1 int

func change(i int) {
	mutex.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1

	if v1 == 10 {
		v1 = 0
		fmt.Print("* ")
	}
	mutex.Unlock()
}

func read() int {
	mutex.Lock()
	a := v1
	mutex.Unlock()
	return a
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	fmt.Printf("%d ", read())

	for i := 0; i< numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf("-> %d", read())
		}(i)
	}

	wg.Wait()
	fmt.Printf("-> %d\n", read())
}