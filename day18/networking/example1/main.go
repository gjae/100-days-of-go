/**
* Page 49 from network programming with Go.pdf
*/
package main

import (
	"fmt"
	"net"
	"os"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s ip-addr\n", os.Args[0])
	}

	name := os.Args[1]

	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid addr")
	} else {
		fmt.Println("The address is ", addr.String())
	}

	os.Exit(0)
}