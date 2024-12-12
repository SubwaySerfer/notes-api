package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes-api/models"
	"notes-api/storage"
	"time"

	"github.com/google/uuid"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {

	note := models.Note{
		ID:        uuid.New().String(),
		Title:     "Title 1",
		Content:   "Content",
		CreatedAt: time.Now().Format(time.RFC3339),
		Author:    "Author",
	}


	err := storage.CreateNote("data/data.json", note)
	if err != nil {
		http.Error(w, "Error creating note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Note created successfully",
		"note":    note,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}

	fmt.Println("Create Note ", note.ID)
	storage.CreateNote("data/data.json", note)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Notes")
	notes, err := storage.LoadNotes("data/data.json")
	if err != nil {
		http.Error(w, "Error loading notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}

}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Note By ID")

	idStr := r.URL.Path[len("/notes/"):]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	notes, err := storage.LoadNotes("data/data.json")
	if err != nil {
		http.Error(w, "Error loading notes", http.StatusInternalServerError)
		return
	}

	var note models.Note
	for _, n := range notes {
		if n.ID == id.String() {
			note = n
			break
		}
	}

	if note.ID == "" {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}


// GET /notes

// Retrieve all notes.
// GET /notes/{id}

// Retrieve a specific note by its ID.
// POST /notes

// Create a new note.
// PUT /notes/{id}

// Update an existing note.
// DELETE /notes/{id}

// Delete a note by its ID.
// ID string `json:"id"`
// Title string `json:"title"`
// Content string `json:"content"`
// CreatedAt string `json:"created_at"`
// Author string `json:"author"`
