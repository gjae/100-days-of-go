package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	handlers "github.com/gjae/recipes-mongo-api/handlers"
)

type Recipe struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Tags         []string           `json:"tags" bson:"tags"`
	Ingredients  []string           `json:"ingredients" bson:"ingredients"`
	Instructions []string           `json:"instructions" bson:"instructions"`
	PublishedAt  time.Time          `json:"publishedat" bson:"publishedAt"`
}

var ctx context.Context
var err error
var client *mongo.Client
var collection *mongo.Collection
var PORT = ":8000"
var redisClient *redis.Client

var recipeHandlers *handlers.RecipeHandler

func init() {
	ctx = context.Background()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	log.Println("Connected to MongoDB")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping(ctx)

	fmt.Println(status)
	recipeHandlers = handlers.NewRecipeHandler(ctx, collection, redisClient)

}

func main() {
	router := gin.Default()
	if len(os.Args) > 1 {
		PORT = ":" + os.Args[1]
	}

	router.GET("/recipes", recipeHandlers.ListRecipesHandler)
	router.POST("/recipes", recipeHandlers.NewRecipeHandler)
	router.PUT("/recipes/:id", recipeHandlers.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipeHandlers.DeleteRecipeHandler)
	router.GET("/recipes/:id", recipeHandlers.FindOneRecipe)
	router.Run(PORT)
	fmt.Println("Connected to mongodb")
}
