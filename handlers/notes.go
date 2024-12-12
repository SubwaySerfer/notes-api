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
	storage.CreateNote("data/data.json", models.Note{ID: "", Title: "Title 1", Content: "Content", CreatedAt: "", Author: "Author"})

	// err = json.NewEncoder(w).Encode(response)
	// if err != nil {
	// 	// Если произошла ошибка при кодировании JSON
	// 	http.Error(w, "Error encoding response", http.StatusInternalServerError)
	// }
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Notes")
	storage.LoadNotes("data/data.json")
}

// ID string `json:"id"`
// Title string `json:"title"`
// Content string `json:"content"`
// CreatedAt string `json:"created_at"`
// Author string `json:"author"`
