package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gjae/recipes-mongo-api/models"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeHandler struct {
	collection  *mongo.Collection
	ctx         context.Context
	redisClient *redis.Client
}

func NewRecipeHandler(ctx context.Context, collection *mongo.Collection, reedisClient *redis.Client) *RecipeHandler {
	return &RecipeHandler{
		collection:  collection,
		ctx:         ctx,
		redisClient: reedisClient,
	}
}

func (handler *RecipeHandler) ListRecipesHandler(c *gin.Context) {
	val, err := handler.redisClient.Get(handler.ctx, "recipes").Result()

	if err == redis.Nil {
		log.Printf("Request to MongoDB")
		curr, err := handler.collection.Find(handler.ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer curr.Close(handler.ctx)

		recipes := make([]models.Recipe, 0)

		for curr.Next(handler.ctx) {
			var recipe models.Recipe
			curr.Decode(&recipe)
			recipes = append(recipes, recipe)
		}

		data, _ := json.Marshal(recipes)
		handler.redisClient.Set(handler.ctx, "recipes", string(data), 0)
		c.JSON(http.StatusOK, recipes)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		log.Printf("Request to redis")
		recipes := make([]models.Recipe, 0)
		json.Unmarshal([]byte(val), &recipes)
		c.JSON(http.StatusOK, recipes)
	}
}

func (handler *RecipeHandler) NewRecipeHandler(c *gin.Context) {
	var recipe models.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	recipe.ID = primitive.NewObjectID()
	recipe.PublishedAt = time.Now()

	_, err := handler.collection.InsertOne(handler.ctx, recipe)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}
	handler.redisClient.Del(handler.ctx, "recipes")
	c.JSON(http.StatusOK, recipe)
}

func (handler *RecipeHandler) UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe models.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectId, _ := primitive.ObjectIDFromHex(id)

	_, err := handler.collection.UpdateOne(handler.ctx, bson.M{
		"_id": objectId,
	}, bson.D{{"$set", bson.D{
		{"name", recipe.Name},
		{"instructions", recipe.Instructions},
		{"ingredients", recipe.Ingredients},
		{"tags", recipe.Tags},
	}}})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	go func(id string, handler *RecipeHandler) {
		handler.redisClient.Del(handler.ctx, "recipes")
		handler.redisClient.Del(handler.ctx, fmt.Sprintf("recipe-%s", id))
	}(id, handler)
	c.JSON(http.StatusOK, gin.H{"message": "Recipe updated successfuly!"})
}

func (handler *RecipeHandler) DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	_, err := handler.collection.DeleteOne(handler.ctx, bson.M{
		"_id": objectId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go func(id string, handler *RecipeHandler) {
		handler.redisClient.Del(handler.ctx, "recipes")
		handler.redisClient.Del(handler.ctx, fmt.Sprintf("recipe-%s", id))
	}(id, handler)

	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been deleted"})
}

func (handler *RecipeHandler) FindOneRecipe(c *gin.Context) {
	id := c.Param("id")
	redisKey := fmt.Sprintf("recipe-%s", id)
	val, err := handler.redisClient.Get(handler.ctx, redisKey).Result()
	var recipe models.Recipe

	if err == redis.Nil {
		log.Printf("Find recipe instance in MongoDB")
		objectId, _ := primitive.ObjectIDFromHex(id)

		object := handler.collection.FindOne(handler.ctx, bson.D{{"_id", objectId}})

		if object.Err() != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			log.Printf(object.Err().Error())
			return
		}
		object.Decode(&recipe)

		data, _ := json.Marshal(recipe)

		handler.redisClient.Set(handler.ctx, fmt.Sprintf("recipe-%s", id), string(data), 0)
		c.JSON(http.StatusOK, recipe)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		log.Printf("Find in redis cache")
		json.Unmarshal([]byte(val), &recipe)
		c.JSON(http.StatusOK, recipe)
	}
}
