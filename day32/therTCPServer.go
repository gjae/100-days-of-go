package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)


func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	SERVER := "localhost:"+args[1]
	s, err := net.ResolveTCPAddr("tcp", SERVER)

	if err != nil {
		fmt.Println(err)
		return 
	}

	l, err := net.ListenTCP("tcp", s)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	buffer := make([]byte, 1024)
	c, err := l.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, err := c.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("> ", string(buffer[0:n-1]), "\n")
		if strings.TrimSpace(string(buffer)) == "STOP" {
			fmt.Println("Exiting server ...")
			return
		}
		c.Write(buffer[0:n-1])
	}
}