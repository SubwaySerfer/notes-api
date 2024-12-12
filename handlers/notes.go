package handlers

import (
	"fmt"
	"net/http"
	"notes-api/models"
	"notes-api/storage"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Note 1")
	storage.CreateNote("data/data.json", models.Note{ID: "", Title: "Title 1", Content: "Content", CreatedAt: "", Author: "Author"})
}

// ID string `json:"id"`
// Title string `json:"title"`
// Content string `json:"content"`
// CreatedAt string `json:"created_at"`
// Author string `json:"author"`
