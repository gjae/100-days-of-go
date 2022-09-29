package main

import "os"

/*
Panic is used for unexpected errors in  normal operation
*/
func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}
}