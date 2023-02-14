package todos

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title   string `json:"title"`
	Text    string `json:"text"`
	BoardID uint   `json:"board_id"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}

func InsertTestTodo(title, text string) *Todo {
	todo := &Todo{Title: title, Text: text}

	return todo
}
