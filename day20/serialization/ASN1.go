package main

import (
	"fmt"
	"encoding/asn1"
)


func main() {
	val := 13

	fmt.Printf("Before Marshal/Unmarshal: %d\n", val)
	mdata, err := asn1.Marshal(val)
	if err != nil {
		panic(err)
	}

	fmt.Printf("After Marshal: %v\n", mdata)
	var n int

	_, err1 := asn1.Unmarshal(mdata, &n)
	if err1 != nil {
		panic(err1)
	}

	fmt.Printf("After Marshal/Unmarshal: %d", n)
}