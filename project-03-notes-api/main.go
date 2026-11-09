package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes []Note

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, "Welcome to Notes API! | Hilal")
}

func notesHandler(w http.ResponseWriter, r *http.Request) {

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

func singleNoteHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	prefix := "/api/v1/notes/"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	idString := strings.TrimPrefix(r.URL.Path, prefix)

	id, err := strconv.Atoi(idString)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	for _, note := range notes {
		if note.ID == id {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(note)
			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/v1/notes", notesHandler)
	http.HandleFunc("/api/v1/notes/", singleNoteHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
