package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Analiza argumentos,. tiene que ser llamada por cuenta propia
	fmt.Println(r.Form) // Imprime informacion en el form en el lado del servidor
	fmt.Println("path", r.URL.Path)
	fmt.Println("Schema ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("Key: ", k)
		fmt.Println("Val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie") // Envia datos al lado del cliente
}

func main() {
	http.HandleFunc("/", sayHelloName)

	err := http.ListenAndServe(":9001", nil)

	if err != nil {
		log.Fatal("Server error: ", err)
	}
}