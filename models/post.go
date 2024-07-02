package models

import "time"

type Post struct {
	Title     string `gorm:"title"`
	Author    string `gorm:"author"`
	Content   string `gorm:"content"`
	ID        uint   `gorm:"primaryKey"`
	CreatedAt time.Time
}
