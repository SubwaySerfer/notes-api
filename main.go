package main

import (
	"log"
	_ "fmt"
	"net/http"
	"notes-api/handlers"
	"notes-api/storage"
	_ "notes-api/db"

	"github.com/gorilla/mux"
)

func main() {

	dbPath := "data/notes.db"
	db, err := storage.EnsureDBExists(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/notes", handlers.CreateNote(db.Conn)).Methods("POST")
	router.HandleFunc("/notes", handlers.GetAllNotes(db.Conn)).Methods("GET")
	router.HandleFunc("/notes/{id}", handlers.GetNoteByID(db.Conn)).Methods("GET")
	router.HandleFunc("/notes/{id}", handlers.UpdateNoteByID(db.Conn)).Methods("PUT")
	router.HandleFunc("/notes/{id}", handlers.DeleteNoteByID(db.Conn)).Methods("DELETE")


	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	if err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
