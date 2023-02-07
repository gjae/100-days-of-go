package main

import "fmt"

// Row matrix is a 1XN matrix (one row by N columns)
func rowMatrix() {
	rowMatrix := [1][3]int{
		{1, 2, 3},
	}

	fmt.Print("Row matrix: ", rowMatrix)
}

// Col matrix is a MX1 matrix (M rows by 1 columns)
func colMatrix() {
	colMatrix := [3][1]int{
		{1},
		{2},
		{3},
	}

	fmt.Print("Column matrix: ", colMatrix)
}

// A lower triangular matrix is when an matrix MXM is filled with zeros above main diaagonal
func lowerTriangularMatrix() {
	m := [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{2, 1, 1},
	}

	fmt.Print("Lower triangular matrix: ", m)
}

// An upper triangular matrix is when a matrix MXM filled with zeros below of main diagonal
func upperTriangularMatrix() {
	m := [3][3]int{
		{1, 2, 3},
		{0, 1, 2},
		{0, 0, 4},
	}

	fmt.Print("Upper triangular matrix: ", m)
}

// Null matrix is when a matrix is filled just with zeros
func nullMatrix() {
	m := [3][3]int{}

	fmt.Print("Null matrix ", m)
}

// Identity matrix is a unit matrix with ones on the main diagonal and
// zeros are elsewhere
func identityMatrix() {
	m := [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	fmt.Print("Identity matrix: ", m)
}

func main() {
	arr := [4][5]int{
		{4, 5, 7, 8, 9},
		{1, 2, 3, 4, 5},
		{9, 10, 11, 12, 14},
		{3, 5, 6, 8, 9},
	}

	var value int = arr[2][3]

	fmt.Print(value)
	rowMatrix()
	fmt.Println()
	colMatrix()
	fmt.Println()
	lowerTriangularMatrix()
	fmt.Println()
	upperTriangularMatrix()
	fmt.Println()
	nullMatrix()
	fmt.Println()
	identityMatrix()
}
