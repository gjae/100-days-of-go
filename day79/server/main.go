package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const SOCKET_PATH = "/tmp/testsocket.sock"

type Server struct {
	Listener net.Listener
}

type Client struct {
	Conn *net.Conn
	IP   net.Addr
}

func NewServer() (*Server, error) {
	os.Remove(SOCKET_PATH)
	listener, err := net.Listen("unix", SOCKET_PATH)
	if err != nil {
		return nil, err
	}

	return &Server{Listener: listener}, nil
}

func HandleConnection(client *Client, ctx *context.Context) {
	data := make(chan string)
	errCh := make(chan error)

	go func(data chan string, errChan chan error, conn *net.Conn) {

		for {
			message := make([]byte, 1024)
			readed, err := (*conn).Read(message)
			if err == io.EOF {
				log.Printf("Client %v - disconnected\n", (*conn).LocalAddr().String())
				errChan <- err
				break
			} else if readed > 0 {
				data <- string(message)
			}
		}

	}(data, errCh, client.Conn)

	for {
		select {
		case <-(*ctx).Done():
			return
		case message := <-data:
			message = strings.ToLower(message)
			message = strings.TrimRight(message, "\n")
			if strings.Contains(message, "exit") {
				fmt.Printf("Client %v say goodbye!\n", client.IP.String())
				return
			}

			fmt.Printf("Client %v -  New message: %s\n", client.IP.String(), message)
		case err := <-errCh:
			log.Printf("Error in socket has been happened: %v", err)
			(*ctx).Deadline()
			return
		}
	}
}

func (l *Server) StartServer(ctx *context.Context) {
	for {
		conn, err := l.Listener.Accept()
		fmt.Print("Waiting for client  message \n")
		if err != nil {
			panic(err)
		}
		client := &Client{Conn: &conn, IP: conn.RemoteAddr()}
		go HandleConnection(client, ctx)
	}
}

func main() {
	server, err := NewServer()
	ctx := context.Background()
	ctx, CancelFunc := context.WithCancel(ctx)
	defer CancelFunc()

	if err != nil {
		log.Fatalf("Error opening the socket: %v", err)
	}
	defer server.Listener.Close()
	server.StartServer(&ctx)

}
