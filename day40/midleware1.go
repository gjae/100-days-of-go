package main

import (
	"fmt"
	"net/http"
	"log"
)

func middleware1(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before call http handler")
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after call real handler")
	})
}

func realHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running real http handler")
	w.Write([]byte("Ok"))
}

func main() {
	originalHandler := http.HandlerFunc(realHandler)

	http.Handle("/", middleware1(originalHandler))

	log.Println(http.ListenAndServe(":8000", nil))
}