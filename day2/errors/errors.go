package main

import (
	"errors"
	"fmt"
)
const NUMERROR int = 42

type argError struct {
	arg int
	prob string
}


func (arg *argError) Error() string {
	return fmt.Sprintf("%d - %s", arg.arg, arg.prob)
}


// An error can be triggered by erros.New function
// errors.New is a basic error, and take an string message as param
func f1(arg int) (int, error) {
	
	if arg == NUMERROR {
		return -1, errors.New("Cant work with 42")
	}

	return arg, nil
}

// Return and pointer error with custon struct error format
func f2(arg int) (int, error) {
	if arg == NUMERROR {
		newError := &argError{prob: "An error has been occurred"}
		return -1, newError
	}

	return arg, nil
}


func main() {
	test1, _ := f1(43)
	test2, _ := f1(42)

	test3, _ := f2(42)

	fmt.Println("test1 worked: ", test1)
	fmt.Println("test2 failed: ", test2)
	fmt.Println("test3 failed: ", test3)
}