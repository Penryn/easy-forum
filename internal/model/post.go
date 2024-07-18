package model

import "time"

type Post struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	UserID  int       `json:"user_id"`
	Time    time.Time `json:"time"`
}
