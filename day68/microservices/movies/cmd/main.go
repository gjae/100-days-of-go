package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"moviesexample.com/pkg/discovery"
	"moviesexample.com/pkg/discovery/consul"

	"moviesexample.com/movies/internal/controller/movie"
	metadatagateway "moviesexample.com/movies/internal/gateway/metadata/http"
	ratinggateway "moviesexample.com/movies/internal/gateway/rating/http"
	httphandler "moviesexample.com/movies/internal/handler/http"
)

const serviceName = "movie"

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
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)

	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)

	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
