package main

import (
	"reflect"
	"fmt"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Flied3 Secret
}


func main() {
	A := Record{"String value", -12.123, Secret {"Mihalis", "Tsoukalos"}}

	r := reflect.ValueOf(A)

	fmt.Println("String value: ", r.String())

	iType  := r.Type()


	fmt.Printf("i Type: %s\n", iType)

	for i := 0; i< r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\t with type: %s ", r.Field(i).Type())
		fmt.Printf("\t and value _%v_", r.Field(i).Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()

		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}
	}
}