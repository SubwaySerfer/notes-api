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


	// // Connect to the database
	// database, err := db.ConnectDB(dbPath)
	// if err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }

	// // If connection is successful, print success message
	// fmt.Println("Successfully connected to the database!")

	// // Example of accessing the GORM connection
	// fmt.Printf("Database connection: %+v\n", database.Conn)


	router := mux.NewRouter()
	router.HandleFunc("/notes", handlers.CreateNote(db.Conn)).Methods("POST")
	router.HandleFunc("/notes", handlers.GetAllNotes(db.Conn)).Methods("GET")
	router.HandleFunc("/notes/{id}", handlers.GetNoteByID(db.Conn)).Methods("GET")
	router.HandleFunc("/notes/{id}", handlers.UpdateNoteByID(db.Conn)).Methods("PUT")
	router.HandleFunc("/notes/{id}", handlers.DeleteNoteByID(db.Conn)).Methods("DELETE")


	// err := http.ListenAndServe(":8080", nil)
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	if err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
