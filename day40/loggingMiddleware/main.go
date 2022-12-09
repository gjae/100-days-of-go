package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"os"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Proccesing request!")

	w.Write([]byte("OK"))
	log.Println("Finished proccesing request")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handle)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000", loggedRouter)
}