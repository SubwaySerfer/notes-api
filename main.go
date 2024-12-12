package main

import (
	"log"
	"net/http"
	"notes-api/handlers"
	"notes-api/storage"
)

func main() {
	storage.EnsureJSONFileExists("data/data.json")

	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost { // Check for POST method
			handlers.CreateNote(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
