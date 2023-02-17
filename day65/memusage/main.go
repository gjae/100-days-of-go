package main

import (
	"log"
	"runtime"
	"syscall"
)

type CMemStats runtime.MemStats

func (mem *CMemStats) AllocbToByte() uint64 {
	return mem.Alloc / 1024 / 1024
}

func (mem *CMemStats) TotalAllocbToB() uint64 {
	return mem.TotalAlloc / 1024 / 1024
}

func main() {
	var men runtime.MemStats
	var info syscall.Sysinfo_t
	_ = syscall.Sysinfo(&info)
	runtime.ReadMemStats(&men)
	mem := CMemStats(men)

	log.Printf("In usage total: \t%dGb\n", mem.Sys/1024/1024)
	log.Printf("Physical mem: \t%.2dGb\n", info.Totalram/1024/1024/1024)
	log.Printf("Free mem: \t\t%dGb", (info.Totalram-(men.Sys*1024))/1024/1024/1024)
}
