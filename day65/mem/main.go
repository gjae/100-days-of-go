package main

import (
	"log"
	"runtime"
)

func main() {
	cpus := runtime.NumCPU()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	log.Printf("Num of CPU in this machine %d", cpus)

	log.Printf("Memory in use: %d\n", mem.Alloc)
	log.Printf("Obtain from system: %d\n", mem.Sys)
	log.Printf("Memory free: %d\n", mem.Frees)
}
