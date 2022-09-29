/*
Learn from  "Build web application with colang"
github https://github.com/astaxie/build-web-application-with-golang/blob/master/es/02.5.md
*/
package main

import "fmt"

type Human struct {
	name string
	age int
}


type Student struct {
	Human // this mark Employee as inherited from Human struct
	university string
}

type Employee struct {
	Human // this mark Employee as inherited from Human struct
	company string
}

// Human struct method
func (h *Human) sayHello() {
	fmt.Printf("Hello! %v \n", h.name)
}

// Method overload for employee struct
func (e *Employee) sayHello() {
	// fmt.Printf("Hello %v!, you are employee at %v", e.Human.name, e.company)
	fmt.Printf("Hello %v!, you are employee at %v", e.name, e.company)
}

func main() {
	studen1 := Student{Human{"Giovanny", 27}, "LUZ/IUTM"}
	employee1 := Employee{Human{"Jaimito", 32}, "GOOGLE"}

	studen1.sayHello()
	employee1.sayHello()
}
