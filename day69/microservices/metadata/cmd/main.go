package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"moviesexample.com/gen"
	"moviesexample.com/metadata/internal/controller/metadata"
	grpchandler "moviesexample.com/metadata/internal/handler/grpc"
	"moviesexample.com/metadata/internal/repository/memory"
	"moviesexample.com/pkg/discovery"
	"moviesexample.com/pkg/discovery/consul"
)

const serviceName = "metadata"

func main() {

	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting the metadata service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)
	log.Println("Starting the movie metadata service")

	repo := memory.New()
	svc := metadata.New(repo)

	h := grpchandler.New(svc)

	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("Failed to lesten: %v", err)
	}

	srv := grpc.NewServer()
	gen.RegisterMetadataServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
