package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"strings"
)


func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MEthod: ", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		//logic part of log in
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //Analizar los parámetros de URL , a continuación, analizar el paquete de respuesta para el cuerpo POST (cuerpo de la solicitud) 
	// Atención: si usted no llama método ParseForm, los siguientes datos no pueden obtenerse del form
	fmt.Println(r.Form) // impresión en el lado del servidor.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // escribir datos de respuesta

}

func main() {
	http.HandleFunc("/", sayHelloName) // setting router rule
	http.HandleFunc("/login", loginHandler)
	err := http.ListenAndServe(":9090", nil)

	if err == nil {
		log.Fatal("Error ", err)
	}
}