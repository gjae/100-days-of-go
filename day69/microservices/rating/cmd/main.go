package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"context"
	"time"

	"google.golang.org/grpc"

	"moviesexample.com/gen"
	"moviesexample.com/pkg/discovery"
	"moviesexample.com/pkg/discovery/consul"

	rating "moviesexample.com/rating/internal/controller"
	grpchandler "moviesexample.com/rating/internal/handler/grpc"
	"moviesexample.com/rating/internal/repository/memory"
)

const serviceName = "rating"

func main() {
	var port int

	log.Println("Starting rating Service")

	repo := memory.New()
	svc := rating.New(repo)
	h := grpchandler.New(svc)

	lis, err := net.Listen("tcp", "localhost:8082")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()

	log.Printf("Starting the metadata service on port %d", port)

	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceIO := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceIO, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceIO, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}

			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceIO, serviceName)

	srv := grpc.NewServer()
	gen.RegisterRatingServiceServer(srv, h)

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
