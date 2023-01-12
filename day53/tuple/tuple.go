package main

import "fmt"

func PowerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func PowerSeriesNamed(a int) (square int, cube int) {
	square = a * a
	cube = square * a

	return
}

func main() {
	var square, cube int

	square, cube = PowerSeries(3)

	fmt.Println("Square ", square, " Cube ", cube)

	square, cube = PowerSeriesNamed(2)

	fmt.Println("Square ", square, " Cube ", cube)
}
