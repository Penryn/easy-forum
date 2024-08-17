package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int            `json:"id"`
	Content   string         `json:"content"`
	UserID    int            `json:"user_id"`
	Time      time.Time      `json:"time"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
