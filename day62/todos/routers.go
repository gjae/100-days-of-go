package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjae/go-todo/commons"
)

func NewTodo(c *gin.Context) {
	var todo *Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := commons.GetCurrentDB()

	db.Create(todo)

	c.JSON(http.StatusCreated, todo)
}

func Todos(c *gin.Context) {
	db, _ := commons.GetCurrentDB()
	var todo []Todo

	db.Where("deleted_at is null").Find(&todo)

	c.JSON(http.StatusOK, todo)

}

func DefineRouter(r *gin.RouterGroup) {
	r.POST("/", NewTodo)
	r.GET("/", Todos)
}
