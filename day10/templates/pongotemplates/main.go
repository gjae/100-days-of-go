package main

import (
	"fmt"
	template "github.com/flosch/pongo2/v6"
)

func main() {
	tpl, err := template.FromString("Hello {{name|capfirst}}")
	if err != nil {
		panic(err)
	}

	out, errt := tpl.Execute(template.Context{"name": "Giovanny"})

	if errt != nil {
		panic(err)
	}

	fmt.Println(out)
}