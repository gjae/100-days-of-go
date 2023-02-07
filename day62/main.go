package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gjae/go-todo/boards"
	"github.com/gjae/go-todo/commons"
	"github.com/gjae/go-todo/config"
	"github.com/gjae/go-todo/todos"
	"gorm.io/gorm"
)

func OpenDatabase(settings config.Settings) (*gorm.DB, error) {
	return commons.GetDb(settings)
}

func DefineAPI() *gin.Engine {
	r := gin.Default()

	boardsApi := r.Group("boards")

	boards.DefineRouter(boardsApi)

	return r

}

func main() {
	settings, _ := config.NewConfig(config.Debug)
	appPort := settings.APIPort
	db, err := OpenDatabase(*settings)
	if err != nil {
		panic(err.Error())
	}

	todos.AutoMigrate(db)
	boards.AutoMigrate(db)

	(DefineAPI()).Run(
		fmt.Sprintf(":%d", appPort),
	)
}
