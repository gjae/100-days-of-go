package main

import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"time"
	"sync"
)

var write = make(chan int)
var read = make(chan int)

// if some function want get the current value
// shoud be readed from read channel
func readVal() int {
	return <-read
}

/**
* When some gorutine want update the variable value
* should call set, or send the new value through write channel
* then the monitor function make the change 
*/
func set(n int) {
	write <- n
}

/**
* This function avoid race conditions
* controlling acces to value Through channels
* of read and writing
*/ 
func monitor() {
	var val int
	for {
		select {
		case x := <- write:
			val = x
			fmt.Printf("%d ", x)
		case read <- val:

		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give a integer!")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	rand.Seed(time.Now().Unix())

	go monitor()
	
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn( 10 * n))
		}()
	}

	wg.Wait()

	fmt.Printf("\nLast value: %d\n", readVal())
}