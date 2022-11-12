package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name	string 	`json:"username"`
	Surname	string 	`json: "surname"`
	Year	int 	`json: "created"`
}

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
	useall := UseAll{Name: "Mike", Surname: "Tsoukalos", Year: 2022}

	// Regular structure
	// Encoding JSON data -> Convert Go Structure to JSON record with fields
	t, err := json.Marshal(&useall)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value: %s\n", t)
	}

	// Decoding JSON data given as a string

	str := `{"username": "M.", "surname": "Ts", "created": 2022}`

	// Convert string into a byte slice

	jsonRecord := []byte(str)

	tmp := UseAll{}
	err = json.Unmarshal(jsonRecord, &tmp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T with value: %v\n", tmp, tmp)
	}
}