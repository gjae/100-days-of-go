package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("testfolder1", 777)
	os.MkdirAll("testfolder1/testa/testb", 777)
	err := os.Remove("testfolder1")

	if err != nil {
		panic(err)
	}

	os.RemoveAll("testfolder1")

	fmt.Println("Ok")
}