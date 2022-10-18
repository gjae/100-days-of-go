package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type User struct {
	Name, Lastname string
}

type RPCHandler struct {

}

func (h *RPCHandler) SayHelloName(payload string, reply *string) error {
	fmt.Printf("Hello from RPC Call: %s", payload)

	*reply = fmt.Sprintf("=== %s === Already was greeted", payload)

	return nil
}


func main() {
	RPCHandlerO := new(RPCHandler)
	rpc.Register(RPCHandlerO)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":3334", nil)

	if err != nil {
		log.Fatalf("Error ", err)
	}
}