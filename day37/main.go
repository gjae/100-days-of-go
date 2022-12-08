package main


import (
	"errors"
	"fmt"
	"os"
	"io"
)

func run(args []string, stdout io.Writer) error {
	if len(args) == 1 {
		return errors.New("No input!")
	}
	fmt.Println("Run function has ben taked a param")
	return nil
}


func main() {
	hasError := run(os.Args, nil)

	if hasError != nil {
		fmt.Println(hasError)
	}
}