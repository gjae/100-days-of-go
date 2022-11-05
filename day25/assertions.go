package main

import (
	"fmt"
)

func returnNumber() interface {} {
	return 12
}

func main() {
	aInt := returnNumber()

	number := aInt.(int)

	number++
	fmt.Println(number)

	// The next statement would fail because tghere 
	// is not type assertion to get the value
	// anInt++
	// the next statement fails but the failure is under
	// control because of the ok bool variable that tells
	// Whether the type assertion is successful or not

	value, ok := aInt.(int64)

	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed")
	}

	// The next statement is successful but
	// dangerous because it does not make sure that
	// the type assertion is successful
	i := aInt.(int)

	fmt.Println("i: ", i)

	_ = aInt.(bool)
}