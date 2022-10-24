package main

import (
	"fmt"
	"os"
	"encoding/asn1"
	"net"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}


func main() {
	service := ":1200"

	addr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err)

	listener, err := net.ListenTCP("tcp4", addr)
	CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Printf("A new connection from %s / %s address\n", conn.RemoteAddr().String(), conn.LocalAddr().String())

		daytime := time.Now()
		// Ignore return network errors.
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close()
	}
}