package main

import (
	"fmt"
	"encoding/json"
)


// Ignoring empty fields in JSON
type NoEmpty struct {
	Name	string	`json:"username"`
	Surname	string 	`json:"surname"`
	Year	int		`json:"created,omitempty"`
}

// Removing private fields nad ignoring emptyfields 
type Password struct {
	Name	string	`json:"username"`
	Surname	string	`json:"surname"`
	Year	int		`json:"created,omitempty"`
	Pass	string	`json:"-"`
}

func main() {
	empty := NoEmpty{Name: "Giovanny", Surname: "Avila"}
	pass := Password{Name: "Giovanny", Surname: "Avila", Pass: "123456"}

	t, err := json.Marshal(&empty)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value: %s\n", t)
	}

	t, err = json.Marshal(pass)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Paass structure value: %s\n", t)
	}

	noEmptyAux := NoEmpty{}
	json1 := `{"username": "USERNAME1", "surname": "Surname1"}`
	err = json.Unmarshal([]byte(json1), &noEmptyAux)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Unmarshal 1: %v\n", noEmptyAux)
	}
}