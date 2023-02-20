package main

import (
	"log"
	"net/http"

	rating "moviesexample.com/rating/internal/controller"
	"moviesexample.com/rating/internal/repository/memory"

	httphandler "moviesexample.com/rating/internal/handler"
)

func main() {
	log.Println("Starting the rating service")

	repo := memory.New()
	ctrl := rating.New(repo)

	h := httphandler.New(ctrl)

	http.Handle("/rating", http.HandlerFunc(h.Handle))

	if err := http.ListenAndServe(":8002", nil); err != nil {
		panic(err)
	}

}
