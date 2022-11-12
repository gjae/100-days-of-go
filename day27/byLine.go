package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	} 

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s", err)
			break
		}

		fmt.Print(line)
	}

	return nil
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("File not found")
		return
	}

	lineByLine(args[1])
}