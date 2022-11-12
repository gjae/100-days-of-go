package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func handleSignal(sig os.Signal) {
	fmt.Println("handleSignal(): ", sig)
}

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs)

	start := time.Now()
	go func() {
		for {
			sig := <- sigs

			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Execution time: ", duration)
			case syscall.SIGUSR1:
				handleSignal(sig)
				os.Exit(0)
			default:
				fmt.Println("Caught: ", sig)
			}
		}
	}()

	for {
		fmt.Println("+")
		time.Sleep(10 * time.Second)
	}
}