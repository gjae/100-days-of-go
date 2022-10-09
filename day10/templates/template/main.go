package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email string
}

func main() {
	person := &Person{UserName: "Giovanny"}
	t := template.New("fieldname example")
	t, _ = t.Parse("Hello {{.UserName}}! {{.email}}")

	t.Execute(os.Stdout, person)
}