package handlers

import (
	"encoding/json"
	_ "fmt"
	"io"
	"net/http"
	"notes-api/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateNote creates a new note in the database
func CreateNote(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.Note
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		note := models.Note{
			ID:        uuid.New().String(),
			Title:     input.Title,
			Content:   input.Content,
			CreatedAt: time.Now(),
			Author:    input.Author,
		}

		if err := db.Create(&note).Error; err != nil {
			http.Error(w, "Failed to create note", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Note created successfully",
			"note":    note,
		})
	}
}

// GetAllNotes fetches all notes from the database
func GetAllNotes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var notes []models.Note
		if err := db.Find(&notes).Error; err != nil {
			http.Error(w, "Failed to load notes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	}
}

// GetNoteByID fetches a note by ID from the database
func GetNoteByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/notes/"):]

		var note models.Note
		if err := db.First(&note, "id = ?", id).Error; err != nil {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(note)
	}
}

// UpdateNoteByID updates a note by ID in the database
func UpdateNoteByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/notes/"):]

		var note models.Note
		if err := db.First(&note, "id = ?", id).Error; err != nil {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}

		var updatedFields struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			Author  string `json:"author"`
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := json.Unmarshal(body, &updatedFields); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if updatedFields.Title != "" {
			note.Title = updatedFields.Title
		}
		if updatedFields.Content != "" {
			note.Content = updatedFields.Content
		}
		if updatedFields.Author != "" {
			note.Author = updatedFields.Author
		}

		if err := db.Save(&note).Error; err != nil {
			http.Error(w, "Failed to update note", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(note)
	}
}

// DeleteNoteByID deletes a note by ID from the database
func DeleteNoteByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/notes/"):]

		if err := db.Delete(&models.Note{}, "id = ?", id).Error; err != nil {
			http.Error(w, "Failed to delete note", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
