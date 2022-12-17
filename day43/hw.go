package main

import "fmt"

func PrintSlice[T any](s []T) {
	for _ ,v  := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}


func main() {
	PrintSlice([]int{1,2,3})
	PrintSlice([]string{"A" , "B", "C"})
	PrintSlice([]float64{1.34, 2.4})
}