package storage

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"time"

	"notes-api/models"
)

func EnsureJSONFileExists(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		err = ioutil.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			log.Fatalf("Failed to create data file: %v", err)
		}
	}
}

func LoadNotes(filename string) ([]models.Note, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var notes []models.Note
	err = json.Unmarshal(data, &notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func SaveNotes(filename string, notes []models.Note) error {
	data, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func CreateNote(filename string, newNote models.Note) error {
	notes, err := LoadNotes(filename)
	if err != nil {
		return err
	}

	newNote.ID = uuid.New().String()
	newNote.CreatedAt = time.Now().Format(time.RFC3339)
	notes = append(notes, newNote)

	return SaveNotes(filename, notes)
}
