package main

import (
	"fmt"
	"os"
)
var version string = "0.0.1-beta"
// Simple change
func main() {
	fmt.Println("Version ", version)
	if path, err := os.Getwd(); err == nil {
		_, err := os.Stat(path + "/.git")
		if os.IsNotExist(err) {
			fmt.Printf("%v dir not exists \n", path)
		} else {
			fmt.Println("Dir ",path + "/.git", " already exists")
		}
		fmt.Fprintf(os.Stdout, "dir: %v", path)
	}
}