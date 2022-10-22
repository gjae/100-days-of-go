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

	udpAddr, err := net.ResolveUDPAddr("udp", service)
	CheckError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	CheckError(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		return
	}

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}