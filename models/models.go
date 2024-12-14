package models

import (
	"time"
)

type Note struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Author    string    `gorm:"type:varchar(100)" json:"author"`
}