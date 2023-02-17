package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func analyzeSentiment(data string, resultChan chan string) {
	defer wg.Done()

	if strings.Contains(strings.ToLower(data), "happy") {
		resultChan <- "positive"
	} else {
		resultChan <- "negative"
	}
}

func main() {
	resultChan := make(chan string)
	inputs := []string{
		"I am so happy today!",
		"I hate this weather",
		"Happy birthday!!",
	}

	for _, input := range inputs {
		wg.Add(1)
		go analyzeSentiment(input, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for i := 0; i < len(inputs); i++ {
		fmt.Println(<-resultChan)
	}
}
