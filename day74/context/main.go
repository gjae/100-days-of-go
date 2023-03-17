package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func SleepContext(_ context.Context, made chan<- bool) {
	fmt.Println("Dome something...")
	k := rand.Intn(100)
	<-time.After(time.Duration(k) * time.Second)
	made <- true
}

func main() {
	ctx := context.Background()
	made := make(chan bool)

	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	ctxTimeout, to := context.WithTimeout(ctx, time.Second*2)
	defer to()
	go SleepContext(ctx, made)

	for {
		select {
		case <-ctxCancel.Done():
			fmt.Println("Cancel")
			return
		case <-ctxTimeout.Done():
			fmt.Printf("Timeout")
			return
		case <-made:
			cancel()
		}
	}
}
