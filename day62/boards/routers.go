package boards

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjae/go-todo/commons"
)

func helloBoardWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "World"})
}

func NewBoard(c *gin.Context) {
	var board *BoardSerializer
	db, _ := commons.GetCurrentDB()

	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boardModel := &Board{Title: board.Title}
	db.Create(boardModel)

	c.JSON(http.StatusCreated, boardModel)

}

func ListBoards(c *gin.Context) {
	db, _ := commons.GetCurrentDB()
	var boards []BoardResponseSerializer

	db.Table("boards").Select("id", "title", "created_at", "active").Scan(&boards)

	c.JSON(http.StatusOK, boards)
}

func DefineRouter(r *gin.RouterGroup) {
	r.GET("boards", helloBoardWorld)
	r.POST("/", NewBoard)
	r.GET("/", ListBoards)
}
