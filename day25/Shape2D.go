package main

import (
	"fmt"
	"math"
)

type Shate2D interface {
	Perimeter() float64
}

type Circle struct {
	R float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	a := Circle{R: 1.5}

	fmt.Printf("R %.2f -> Perimeter %.3f \n", a.R, a.Perimeter())

	// THis is like a instanceof Shape2D or isinstance(a, Shape2D)
	_, ok := interface{}(a).(Shate2D)

	if ok {
		fmt.Println("a Is a Shape2D")
	}
}