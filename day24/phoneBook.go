package main

import (
	"fmt"
	"os"
	"path"
	"math/rand"
	"strconv"
)

type Entry struct {
	Name string
	Suername string
	Tel string
}

const MIN = 0
const MAX = 94

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Tel == key {
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

func random(min, max int) int {
    return rand.Intn(max-min) + min
}

func getString(len int64) string {
    temp := ""
    startChar := "!"
    var i int64 = 1
    for {
        myRand := random(MIN, MAX)
        newChar := string(startChar[0] + byte(myRand))
        temp = temp + newChar
        if i == len {
            break
        }
        i++
    }
    return temp
}


func populate(n int, s []Entry) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100, 198))

		data = append(data, Entry{name, surname, n})
	}
}

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage : %s search|list <arguments>\n", exe)
		return 
	}
	populate(8, data)

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