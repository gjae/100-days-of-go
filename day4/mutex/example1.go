package main

import (
	"fmt"
	"sync"
	"time"
	"errors"
)

/**
Mutex can be used with gorutines or only calling method
*/
type SafeCounter struct {
	mu sync.Mutex 
	v map[string]int
}

func (s *SafeCounter) Inc(key string, value int) {
	// Lock struct
	s.mu.Lock()
	s.v[key] += value

	/**
	* unlock struct
	*/
	s.mu.Unlock()
}

func (s *SafeCounter) Value(key string) (int, error) {
	/* 
	We can also use defer to ensure the mutex will be unlocked as in the Value method.
	reminding that *defer keyword* run function after all function body 
	has been executed
	*/
	defer s.mu.Unlock()
	
	s.mu.Lock()
	value, isPresent := s.v[key]	

	if !isPresent {
		return -1, errors.New("Key not present")
	}

	return value, nil
}

func main() {
	data := make(map[string]int)
	data["money"] = 0
	mutex := SafeCounter{v: data}

	for i := 0; i < 5; i++ {
		go mutex.Inc("money", 30)
	}

	time.Sleep(time.Second)
	fmt.Println(mutex.v)
}	