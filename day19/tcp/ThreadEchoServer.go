package main


import (
	"fmt"
	"net"
	"os"
)



func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])

		if err != nil {
			return 
		}

		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])

		if err2 != nil {
			return
		}
	}
}

func main() {
	service := ":1201"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
		conn.Close()
	}
}