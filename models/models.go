package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint      `gorm:"primaryKey"`
	Username       string    `gorm:"unique;not null"`
	HashedPassword string    `gorm:"not null"`
	Role           string    `gorm:"type:text;default:'user';not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

type Note struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UserID    string    `gorm:"type:uuid;"`
}

type Database struct {
	Conn *gorm.DB
}