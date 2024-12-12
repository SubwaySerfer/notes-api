package main

import (
	"log"
	"net/http"
	"notes-api/handlers"
	"notes-api/storage"
)

// @title Notes API
// @version 1.0
// @description API for managing notes
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /

// Ensure the JSON file for storage exists
func main() {
	storage.EnsureJSONFileExists("data/data.json")

	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost { // Check for POST method
			handlers.CreateNote(w, r)
		} else if r.Method == http.MethodGet { // Check for GET method
			handlers.GetAllNotes(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	 // @Summary Get a single note by ID
    // @Description Get a note by its unique ID
    // @Produce  json
    // @Param id path int true "Note ID"
    // @Success 200 {object} handlers.Note
    // @Failure 404 {string} string "Note not found"
    // @Router /notes/{id} [get]

	http.HandleFunc("/notes/", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handlers.GetNoteByID(w, r)
    case http.MethodPut:
        handlers.UpdateNoteByID(w, r)
		case http.MethodDelete:
				handlers.DeleteNoteByID(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
})


	log.Println("Server started on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
