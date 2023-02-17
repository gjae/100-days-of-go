package main

import (
	"log"

	"github.com/shirou/gopsutil/v3/mem"
)

func main() {

	cmem, _ := mem.VirtualMemory()

	log.Print(cmem)

}
