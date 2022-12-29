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

	handlers "github.com/gjae/recipes-auth/handlers"
	middlewares "github.com/gjae/recipes-auth/middlewares"
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
var authHandler *handlers.AuthHandler

var recipeHandlers *handlers.RecipeHandler

func init() {
	// fileData, _ := os.ReadFile("recipes.json")
	// recipes := make([]interface{}, 0)
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
	// json.Unmarshal(fileData, &recipes)
	// _, err := collection.InsertMany(ctx, recipes)

	/* if err != nil {
		panic(err)
	}*/
	fmt.Println(status)
	recipeHandlers = handlers.NewRecipeHandler(ctx, collection, redisClient)
	authHandler = &handlers.AuthHandler{}

}

func main() {
	router := gin.Default()
	if len(os.Args) > 1 {
		PORT = ":" + os.Args[1]
	}

	router.POST("/signin", authHandler.SignInHandler)
	router.GET("/recipes", recipeHandlers.ListRecipesHandler)
	authorization := router.Group("/")

	authorization.Use(middlewares.JWTAuthMiddleware())
	{
		authorization.POST("/recipes", recipeHandlers.NewRecipeHandler)
		authorization.PUT("/recipes/:id", recipeHandlers.UpdateRecipeHandler)
		authorization.DELETE("/recipes/:id", recipeHandlers.DeleteRecipeHandler)
		authorization.GET("/recipes/:id", recipeHandlers.FindOneRecipe)
	}

	router.Run(PORT)
	fmt.Println("Connected to mongodb")
}
