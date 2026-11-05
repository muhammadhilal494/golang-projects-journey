package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes []Note

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Notes API! | Hilal")
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "This is the notes endpoint!")

	// fmt.Fprintln(w, "HTTP Method:", r.Method)
	// fmt.Fprintln(w, "URL Path:", r.URL.Path)

	switch r.Method {
	case http.MethodGet:

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(notes)

	case http.MethodPost:

		var newNote Note

		err := json.NewDecoder(r.Body).Decode(&newNote)

		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		newNote.ID = len(notes) + 1

		notes = append(notes, newNote)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newNote)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")

	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/v1/notes", notesHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
