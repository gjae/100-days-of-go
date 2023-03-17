package main

import (
	"fmt"
	"io"

	"os/exec"
)

func main() {

	cmd := exec.Command("cat")

	stdin, _ := cmd.StdinPipe()

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "Hello world")
		io.WriteString(stdin, " Another message")
	}()

	cmd.Wait()

	out, _ := cmd.CombinedOutput()

	fmt.Printf(">> %s\n", out)
}
