package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is: "
	fmt.Fprintf(w, "<h1 align='center'>%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align='center'>%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	port := ":8001"
	arguments := os.Args

	if len(arguments) != 1 {
		port = ":"+arguments[1]
	}

	fmt.Println("Using port: ", port)
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/time", timeHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}


}