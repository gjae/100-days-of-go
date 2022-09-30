package main

import (
	"fmt"
	"math"
	"reflect"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Geometry interface has been declared in interfaces.go file
// measure receive some object (struct typed) that implements
// Geometry interface
func measure(c Geometry) {
	fmt.Println(c.area())
	fmt.Println(c.perim())
	fmt.Println(reflect.TypeOf(c))
}

func main() {
	c := circle{2.0}
	measure(c)
}