package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

var channelApp *amqp.Channel

type Request struct {
	URL string `json:"url"`
}

func init() {
	amqpConnection, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Fatal(err)
	}

	channelApp, _ = amqpConnection.Channel()
}

func ParseHandler(c *gin.Context) {
	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Print(request)

	data, _ := json.Marshal(request)

	err := channelApp.Publish(
		"",
		os.Getenv("RABBITMQ_QUEUE"),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(data),
		})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func main() {
	router := gin.Default()

	router.POST("/parse", ParseHandler)
	router.Run(":5000")
}
