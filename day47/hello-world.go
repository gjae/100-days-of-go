package main

import "github.com/gin-gonic/gin"

func IndexHandler(c *gin.Context) {
	username := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + username,
	})
}

func main() {
	router := gin.Default()

	router.GET("/:name", IndexHandler)

	router.Run(":8000")
}
