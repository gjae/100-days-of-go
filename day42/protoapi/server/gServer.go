package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
	"github.com/gjae/protoapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RandomServer struct {
	protoapi.UnimplementedRandomServer
}

var min = 0
var max = 100

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Extra function for creating secure random numbers
//
// func randomSecure(min, max int) int {
// 	v, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
// 	if err != nil {
// 		fmt.Println(err)
// 		return min
// 	}
// 	fmt.Println("**", v, min, max)

// 	return min + int(v.Uint64())
// }

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		// For getting valid ASCII characters
		myRand := random(0, 94)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}


func (RandomServer) GetDate(ctx context.Context, r *protoapi.RequestDateTime) (*protoapi.DateTime, error) {
	currentTime := time.Now()
	response := &protoapi.DateTime{
		Value: currentTime.String(),
	}

	return response, nil
}

func (RandomServer) GetRandom(ctx context.Context, r *protoapi.RandomParams) (*protoapi.RandomInt, error) {
	rand.Seed(r.GetSeed())

	place := r.GetPlace()

	temp := random(min, max)
	for {
		place--
		if place <= 0 {
			break
		}
		temp = random(min, max)
	}

	response := &protoapi.RandomInt{
		Value: int64(temp),
	}

	return response, nil
} 

func (RandomServer) GetRandomPass(ctx context.Context, r *protoapi.RequestPass) (*protoapi.RandomPass, error) {
	rand.Seed(r.GetSeed())

	temp := getString(r.GetLength())

	response := &protoapi.RandomPass{
		Password: temp,
	}
	return response, nil
}

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port : ", port)
	} else {
		port = os.Args[1]
	}

	server := grpc.NewServer()

	var randomServer RandomServer
	protoapi.RegisterRandomServer(server, randomServer)

	reflection.Register(server)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Serving Requests ...")
	server.Serve(listen)
}
