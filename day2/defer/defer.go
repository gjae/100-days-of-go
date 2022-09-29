package main

import (
	"fmt"
	"os"
)


func main() {
	file := createFile("/tmp/defer.txt")

	// Defer its used for run function at final of all program
	// After all statements has been ran
	// here use defer for ensure that file pointer is closed at final of program 
	// Observe that closeFile is called before writeFile , but "defer" make call  to closeFile after writeFIle
	defer closeFile(file)
	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("Creating file ...")

	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	return f
}


func closeFile(file *os.File) {
	fmt.Println("Closing file ...")
	err := file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(file *os.File) {
	fmt.Println("Writing file ...")

	// Write to file (F[ile]println(*f, "content"))
	fmt.Fprintln(file, "data")
}