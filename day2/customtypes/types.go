package main

import "fmt"

// Creting a custom types
// Like typedef in C/C++
type skills []string

type User struct {
	name string
	age int
	skills // use custom type as struct field
}

// Declare struct print method
// For show name and skills for stdout
func (u User) print() {
	fmt.Printf("%v: %v", u.name, u.skills)
}

func main() {
	// Declaring a variable as skills type
	userSkills := skills{"PHP", "Java", "Go"}
	
	fmt.Println(userSkills)
	userSkills = append(userSkills, "Python")
	u := User{"Giovanny", 27, userSkills}

	u.print()
}