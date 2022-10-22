package main

import (
	"fmt"
	"os"
	"net"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s dottet-ip-addr \n", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Println("Invalid Address")
		os.Exit(1)
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println("Address is ", addr.String(),
		"\nDefault mask length is ", bits,
		"\nLeading ones count is ", ones,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
	os.Exit(0)
}