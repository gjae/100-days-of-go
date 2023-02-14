package boards

import "time"

type BoardSerializer struct {
	Title string `json:"title"`
}

type BoardResponseSerializer struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

type TodoSerializer struct {
	ID        uint
	Board     Board
	Title     string `gorm:"title" json:"title"`
	Text      string `gorm:"text" json:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
