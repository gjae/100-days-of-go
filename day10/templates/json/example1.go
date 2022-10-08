package main

import (
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Server struct {
	ServerName string
	ServerIP string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	file, err := os.Open("test.json")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	v := Serverslice{}
	json.Unmarshal([]byte(content), &v)

	fmt.Println(v)

	// Parse json to interface
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f map[string]interface{}
	err = json.Unmarshal(b, &f)


	for k, v := range f {
		switch vv := v.(type) {
		case int :
			fmt.Println(k, " Is integer type ", vv)
		case string:
			fmt.Println(k, " Is string type ", vv)
		case float64:
			fmt.Println(k, " Is float64 type ", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println("Undefined")
		}
	}
}