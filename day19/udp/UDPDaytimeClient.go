package main

import (
	"fmt"
	"os"
	"net"
)


func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {

	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp", service)
	CheckError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	CheckError(err)

	_, err = conn.Write([]byte("anything"))
	CheckError(err)

	var buf [512]byte

	n, err := conn.Read(buf[0:])
	CheckError(err)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

