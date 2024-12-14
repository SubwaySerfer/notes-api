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

	// Migrate the schema for the "Note" model
	err = db.AutoMigrate(&models.Note{})
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
func (d *Database) LoadNotes() ([]models.Note, error) {
	var notes []models.Note
	if err := d.Conn.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}



// package storage

// import (
// 	"encoding/json"
// 	"log"
// 	"os"

// 	"notes-api/models"
// )

// func EnsureJSONFileExists(filename string) {
// 	_, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		err = os.WriteFile(filename, []byte("[]"), 0644)
// 		if err != nil {
// 			log.Fatalf("Failed to create data file: %v", err)
// 		}
// 	}
// }

// func LoadNotes(filename string) ([]models.Note, error) {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var notes []models.Note
// 	err = json.Unmarshal(data, &notes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return notes, nil
// }

// func SaveNotes(filename string, notes []models.Note) error {
// 	data, err := json.Marshal(notes)
// 	if err != nil {
// 		return err
// 	}

// 	return os.WriteFile(filename, data, 0644)
// }

// func CreateNote(filename string, newNote models.Note) error {
// 	notes, err := LoadNotes(filename)
// 	if err != nil {
// 		return err
// 	}

// 	notes = append(notes, newNote)

// 	return SaveNotes(filename, notes)
// }
