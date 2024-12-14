package storage

import (
	_ "log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"notes-api/models"
)

type Database struct {
	Conn *gorm.DB
}

// EnsureDBExists initializes the database and ensures the "notes" table is created
func EnsureDBExists(dbPath string) (*Database, error) {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate schemas for both models
	err = db.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}


// CreateNote adds a new note to the database
func (d *Database) CreateNote(note models.Note) error {
	if err := d.Conn.Create(&note).Error; err != nil {
		return err
	}
	return nil
}

// LoadNotes fetches all notes from the database
func (d *Database) LoadNotesByUser(userID uint) ([]models.Note, error) {
	var notes []models.Note
	if err := d.Conn.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}
