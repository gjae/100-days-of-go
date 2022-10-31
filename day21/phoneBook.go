package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name string
	Suername string
	Tel string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Suername == key {
			return &data[i]
		}
	}

	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage : %s search|list <arguments>\n", exe)
		return 
	}

	data = append(data, Entry{"Mihails", "Tsoukalos", "11111111"})
	data = append(data, Entry{"Mary", "Doe", "122222"})
	data = append(data, Entry{"John", "Black", "33312220"})

	// Differentiate between the commands

	switch arguments[1] {
	// the search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return 
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not founds: ", arguments[2])
			return
		}
		fmt.Println(*result)
	// the list command
	case "list":
		list()
	// Response to anything is not ma match
	default:
		fmt.Println("Not a valid option")
		
	}
}