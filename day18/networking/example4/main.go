package main

import (
	"fmt"
	"os"
	"net"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s hostname", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}

	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", name)

	if err != nil {
		fmt.Println("REsolution error, ", err.Error())
	}

	fmt.Println("Resolved address is ", addr.String())
}