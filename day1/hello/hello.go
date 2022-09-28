package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	// Get greeting message and print it
	message, err := greetings.Hello("Giovanny")
	if err  != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// Request greeting messages for the names
	names := []string{"Jaimito", "Pepito", "Perendejo"}
	messages, err := greetings.Hellos(names)

	// Check for Hellos error returned
	if err != nil {
		log.Fatal(err)
	}

	// if greetings.Hellos not return error, then 
	// print to console all messages
	// With this for, i variable get map key and value get the map value
	for i, value := range messages {
		fmt.Printf("%v: %v\n", i, value)
	}
}