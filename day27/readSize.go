package main

import (
	"fmt"
	"os"
	"io"
	"strconv"
)


func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)

	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer[0:n]
}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Please run <program> size filepath")
		return 
	}

	f, err := os.Open(args[2])
	if err != nil {
		fmt.Print(err)
		return
	}

	defer f.Close()
	
	var size int

	size, err = strconv.Atoi(args[1])

	data := readSize(f, size)
	fmt.Println(string(data))
}