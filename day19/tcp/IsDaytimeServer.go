package main

import (
	"fmt"
	"os"
	"net"
	"time"
)


func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)

	CheckError(err)

	listen, err := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listen.Accept()

		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))

		conn.Close()
	}
}