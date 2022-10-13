package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func HandlerMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HandlerMain)

	http.ListenAndServe(":9000", r)
}