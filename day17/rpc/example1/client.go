package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "serve")
	}

	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":3333")

	// Synchronous call
	arith := Args{4, 3}
	var reply int

	err = client.Call("Arith.Multiply", arith, &reply)

	if err != nil {
		log.Fatalf("arith error; ", err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", arith.A, arith.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", arith, &quot)

	if err != nil {
		log.Fatalf("Arith error: ", err)
	}

	fmt.Printf("Arith: %d/%d=%d reminder %d\n", arith.A, arith.B, quot.Quo, quot.Rem)
}