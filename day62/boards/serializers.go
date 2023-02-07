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
