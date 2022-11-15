package main

import (
	"fmt"
	"runtime"
)


func main() {
	fmt.Print("You arre using ", runtime.Compiler, " ")
	fmt.Println("On a ", runtime.GOARCH," Machine")
	fmt.Println("Using Go version ", runtime.Version())

	fmt.Printf("COMAXPROCS %d\n", runtime.GOMAXPROCS(0))
}