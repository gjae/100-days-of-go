package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Request struct {
	URL string `json:"url"`
}

type Feed struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Link struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Thumbnail struct {
		URL string `xml:"url,attr"`
	} `xml:"thumbnail"`
	Title string `xml:"title"`
}

var client *mongo.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
}

func GetFeedEntries(url string) ([]Entry, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var feed Feed
	value, _ := io.ReadAll(resp.Body)
	xml.Unmarshal(value, &feed)

	return feed.Entries, nil
}

var (
	RABBITMQ_URI     = os.Getenv("RABBITMQ_URI")
	RABBITMQ_QUEUE   = os.Getenv("RABBITMQ_QUEUE")
	MONGO_DATABASE   = os.Getenv("MONGO_DATABASE")
	MONGO_COLLECTION = "recipes"
)

func main() {
	amqpConnection, err := amqp.Dial(RABBITMQ_URI)

	if err != nil {
		log.Fatal(err)
	}
	defer amqpConnection.Close()

	channelApp, _ := amqpConnection.Channel()
	defer channelApp.Close()

	forever := make(chan bool)

	msg, err := channelApp.Consume(
		RABBITMQ_QUEUE,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msg {
			var request Request
			json.Unmarshal([]byte(d.Body), &request)
			entries, _ := GetFeedEntries(request.URL)
			collection := client.Database(MONGO_DATABASE).Collection(MONGO_COLLECTION)
			for _, entry := range entries[2:] {
				collection.InsertOne(ctx, bson.M{
					"title":     entry.Title,
					"thumbnail": entry.Thumbnail.URL,
					"url":       entry.Link.Href,
				})
			}
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
