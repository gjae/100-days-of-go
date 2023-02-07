package todos

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title   string
	Text    string
	BoardID uint
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}

func InsertTestTodo(title, text string) *Todo {
	todo := &Todo{Title: title, Text: text}

	return todo
}
