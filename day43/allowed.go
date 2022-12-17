package main

import "fmt"

func Same[T comparable](a, b T) bool {
	return a == b
}

func main() {
	fmt.Println("4 = 3 is ", Same(4, 3))
	fmt.Println("aa = aa is ", Same("aa", "aa"))
	fmt.Println("4.1 = 4.15", Same(4.1, 4.15))
	fmt.Println("[]int{1,2,3} = []int{1,2 ,3}", Same([]int{1,2,3}, []int{1,2,3}))
}