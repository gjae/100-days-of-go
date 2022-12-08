package main

import (
	"io"
	"log"
	"net/http"
	"time"
)


func HealthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/health", HealthCheck)

	log.Print(http.ListenAndServe(":8000", nil))
}