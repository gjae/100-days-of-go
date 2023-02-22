package main

import (
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Printf("Child process ID %d\n", os.Getpid())
	time.Sleep(10 * time.Second)
}
