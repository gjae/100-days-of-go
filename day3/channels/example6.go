package main

import (
	"fmt"
	"time"
)

func worker(worker int, job <- chan int, results chan <- int) {
	for j := range job {
		fmt.Println("Worker ", worker,  "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker ", worker, "Has been finished job ", j)
		results <- j
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < numJobs; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a<=numJobs; a++ {
		<- results
	}
}