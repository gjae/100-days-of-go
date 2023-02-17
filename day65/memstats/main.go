package main

import (
	"log"
	"runtime"
)

func main() {
	runtime.GC()
	var p []runtime.MemProfileRecord
	n, ok := runtime.MemProfile(nil, false)

	for {
		p = make([]runtime.MemProfileRecord, n+50)
		n, ok = runtime.MemProfile(p, false)

		if ok {
			p = p[0:n]
			break
		}
	}

	var total runtime.MemProfileRecord
	for i := range p {
		r := &p[i]
		total.AllocBytes += r.AllocBytes
		total.AllocObjects += r.AllocObjects
		total.FreeBytes += r.FreeBytes
		total.FreeObjects += r.FreeObjects
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Printf("%d in use objects (%d in use bytes) | Alloc: %d TotalAlloc: %d", total.InUseObjects(), total.InUseBytes(), m.Alloc, m.TotalAlloc)
}
