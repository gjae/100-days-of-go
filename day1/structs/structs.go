package main

import "fmt"
import "github.com/spf13/cobra"

type person struct {
	name string
	age int
}

type rect struct {
	width, height int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

// perim and area belongs to rect struct
// as struct methods
func (r *rect) perim() int {
	return r.width * r.height
}

func (r *rect) area() int {
	return 2*r.width + 2*r.height
}


func main() {
	// Severals ways for declare structs
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})

	// Age take default zero value
	fmt.Println(person{name: "Fred"})

	// With pointers
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jhon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp)
	fmt.Println(s)


	// Calling struct methods

	r := rect{width: 10, height: 3}
	fmt.Println(r.perim())
	fmt.Println(r.area())
}