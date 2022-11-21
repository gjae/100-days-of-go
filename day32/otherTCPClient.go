package main

import (
	"fmt"
	"bufio"
	"net"
	"os"
	"strings"
)


func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please provide host:port.")
		return 
	}

	connect := arguments[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := net.DialTCP("tcp4", nil, tcpAddr)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: "+message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client existing...")
			c.Close()
			return
		}
	}
}