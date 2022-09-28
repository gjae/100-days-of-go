package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the name person
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Greeting to multiples names
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)

	// loop through the received slice of names, calling
	// the hello function
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// init sets initials values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random
func randomFormat() string {
	// A slice of messages formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v!, Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}