package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"moviesexample.com/metadata/internal/controller/metadata"
	httpHandler "moviesexample.com/metadata/internal/handler/http"
	"moviesexample.com/metadata/internal/repository/memory"
	"moviesexample.com/pkg/discovery"
	"moviesexample.com/pkg/discovery/consul"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
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

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
