package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type InputEvent struct {
	Link string `json:"link"`
	Key  string `json:"key"`
}

type Response struct {
	Link string `json:"link"`
	Key  string `json:"key"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event InputEvent) (Response, error) {
	fmt.Println(event)
	return Response{
		Link: event.Link,
		Key:  event.Key,
	}, nil
}
