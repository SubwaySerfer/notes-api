package db

import (
	"fmt"
	"os"

	"notes-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database struct with GORM connection
type Database struct {
	Conn *gorm.DB
}

// ConnectDB opens or creates a SQLite database
func ConnectDB(dbPath string) (*Database, error) {
	// Check if database file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Println("Database file does not exist. Creating a new one...")
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create database file: %w", err)
		}
		file.Close() // Close the file after creation
	}

	// Connect to the database
	conn, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Database{Conn: conn}, nil
}

func (d *Database) CreateUser(user models.User) error {
	if err := d.Conn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}