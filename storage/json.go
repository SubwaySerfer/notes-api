package storage

import (
	"encoding/json"
	"log"
	"os"

	"notes-api/models"
)

func EnsureJSONFileExists(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		err = os.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			log.Fatalf("Failed to create data file: %v", err)
		}
	}
}

func LoadNotes(filename string) ([]models.Note, error) {
	data, err := os.ReadFile(filename)
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

	return os.WriteFile(filename, data, 0644)
}

func CreateNote(filename string, newNote models.Note) error {
	notes, err := LoadNotes(filename)
	if err != nil {
		return err
	}

	notes = append(notes, newNote)

	return SaveNotes(filename, notes)
}
