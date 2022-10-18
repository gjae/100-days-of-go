package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"strings"
)

func main() {
	var username string
	var host string
	var port string
	flag.StringVar(&username, "name", "", "Nombre a saludar")
	flag.StringVar(&host, "host", "localhost", "IP/Host del canal RPC")
	flag.StringVar(&port, "port", "3334", "Puerto TCP del canal RPC")
	flag.Parse()

	host = fmt.Sprintf("%s:%s", host, port)

	client, err := rpc.DialHTTP("tcp", host)

	if err != nil {
		panic(err)
	}

	var user string

	username = strings.ToUpper(strings.ToLower(username))
	err = client.Call("RPCHandler.SayHelloName", username, &user)
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", user)
}