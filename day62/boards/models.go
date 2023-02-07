package boards

import (
	"github.com/gjae/go-todo/todos"
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Title   string
	Active  bool `gorm:"default:1"`
	Deleted gorm.DeletedAt
	Todos   []todos.Todo
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Board{})
}
