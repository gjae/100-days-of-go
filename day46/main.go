package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"github.com/gjae/go-mongo1/database"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB successfully")
	collection := client.Database("appDB").Collection("movies")

	darkNight := database.Movie{
		Name: "The DarK Knight",
		Year: "2008",
		Directors: []string{"Cristopher Nolan"},
		Writers: []string{"Jonathan Nolan", "Cristopher Nolan"},
		BoxOffice: database.BoxOffice{
			Budget: 185100000,
			Gross: 533316061,
		},
	}

	_, err = collection.InsertOne(context.TODO(), darkNight)

	if err != nil {
		log.Fatal(err)
	}

	// Querying data

	queryResult := &database.Movie{}

	filter := bson.M{"boxOffice.budget": bson.M{"$gt": 15000000}}

	result := collection.FindOne(context.TODO(), filter)
	err = result.Decode(queryResult)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie: ", queryResult)

	err = client.Disconnect(context.TODO())
	fmt.Println("Disconnected from MongoDB")
}