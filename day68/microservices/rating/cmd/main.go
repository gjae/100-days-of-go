package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"context"
	"time"

	"moviesexample.com/pkg/discovery"
	"moviesexample.com/pkg/discovery/consul"

	rating "moviesexample.com/rating/internal/controller"
	"moviesexample.com/rating/internal/repository/memory"

	httphandler "moviesexample.com/rating/internal/handler"
)

const serviceName = "rating"

func main() {
	var port int
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
	repo := memory.New()
	ctrl := rating.New(repo)

	h := httphandler.New(ctrl)

	http.Handle("/rating", http.HandlerFunc(h.Handle))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}

}
