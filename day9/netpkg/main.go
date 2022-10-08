package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()

	if err != nil {
		panic(err)
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %s \n", i.Name)
	}
}