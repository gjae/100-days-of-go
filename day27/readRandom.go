package main

import (
	"os"
	"encoding/binary"
	"fmt"
)

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	var seed int64

	binary.Read(f, binary.LittleEndian, &seed)

	fmt.Println("Seed: ", seed)
}