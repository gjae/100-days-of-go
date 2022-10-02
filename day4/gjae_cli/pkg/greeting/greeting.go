package greeting

import (
	"fmt"
	"time"
	"math/rand"
)

type Name string
var r *rand.Rand = nil

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
}

func getRandonGreeetingMessage() (message string) {
	greeting := []string{"Hey %v!", "Hello %v!", "Nice to see you %v!", "Good day Sr %v!"}
	message = greeting[r.Intn(len(greeting))]

	return message
}

func (name Name) SayHello(isRandom bool) {
	message := "Hello %v"
	if isRandom {
		message = getRandonGreeetingMessage()
	}
	fmt.Printf(message, name)
}