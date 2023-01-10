package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID          string       `json:"id,omitempty" `
	Name        string       `json:"name"`
	Picture     string       `json:"imageURL"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []string     `json:"steps"`
}

type Ingredient struct {
	Quantity string `json:"quantity"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)

	file, err := os.ReadFile("recipes.json")

	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &recipes)

	for i, _ := range recipes {
		recipes[i].ID = xid.New().String()
	}

}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"recipes": recipes,
	})
}

func RecipeHandler(c *gin.Context) {
	id := c.Param("id")

	for _, r := range recipes {
		if r.ID == id {
			c.HTML(http.StatusOK, "recipe.html", gin.H{
				"recipe": r,
			})

			return
		}
	}

	c.File("404.html")
}

func main() {
	router := gin.Default()
	router.Static("assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			log.Printf("Its a custom middleware")
			c.Next()
		}
	}())

	router.GET("/", IndexHandler)
	router.GET("/recipes/:id", RecipeHandler)

	router.Run(":8000")
}
