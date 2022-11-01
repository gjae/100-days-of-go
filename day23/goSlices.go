package main

import "fmt"

func main() {
	// create an empty slice
	aSlice := []float64{}

	// Both length and capacity area 0 because aSlice is empty
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	// Add elements to a slice
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)

	fmt.Println(aSlice, "With len", len(aSlice))

	// A slice with length 4
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4

	// Now you will need to use append
	t = append(t, -5)
	fmt.Println(t)

	// A 2Dslice
	// You can have as many dimensions as need
	twoD := [][]int{{1,2,3}, {4,5,6}}
	// Visiting all elements of  2D slice
	// With double for loop
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}


	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1,2,3,4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)
}