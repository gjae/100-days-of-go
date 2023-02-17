package main

import (
	"log"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()

	log.Printf("Total CPU: %d", cpu)
}
