package main

import (
	tpl "github.com/flosch/pongo2/v6"
	"net/http"
)

var tplExample = tpl.Must(tpl.FromFile("example.html"))


func examplePongo(w http.ResponseWriter, r *http.Request) {
	err := tplExample.ExecuteWriter(tpl.Context{"query": r.FormValue("query")}, w)

	if err != nil {
		panic(err)
	}
}

func main() {
    http.HandleFunc("/", examplePongo)
    http.ListenAndServe(":8080", nil)
}