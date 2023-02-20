package main

import (
	"log"
	"net/http"

	"moviesexample.com/movies/internal/controller/movie"
	metadatagateway "moviesexample.com/movies/internal/gateway/metadata/http"
	ratinggateway "moviesexample.com/movies/internal/gateway/rating/http"
	httphandler "moviesexample.com/movies/internal/handler/http"
)

func main() {
	log.Println("Starting the movie service")

	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")

	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)

	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))

	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
