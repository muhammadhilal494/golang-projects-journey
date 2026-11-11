package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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

		if !validateNote(&newNote) {
			http.Error(w, "Title and content cannot be empty", http.StatusBadRequest)
			return
		}

		newNote.ID = nextNoteID
		nextNoteID++

		notes = append(notes, newNote)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newNote)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func singleNoteHandler(w http.ResponseWriter, r *http.Request) {

	prefix := "/api/v1/notes/"

	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idString := strings.TrimPrefix(r.URL.Path, prefix)

	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet, http.MethodPut, http.MethodDelete:
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	for i, note := range notes {
		if note.ID == id {

			switch r.Method {

			case http.MethodGet:

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(note)
				return

			case http.MethodPut:

				var updatedNote Note

				err := json.NewDecoder(r.Body).Decode(&updatedNote)

				if err != nil {
					http.Error(w, "Invalid JSON data", http.StatusBadRequest)
					return
				}

				if !validateNote(&updatedNote) {
					http.Error(w, "Title and content cannot be empty", http.StatusBadRequest)
					return
				}

				updatedNote.ID = note.ID
				notes[i] = updatedNote

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(updatedNote)
				return

			case http.MethodDelete:

				notes = append(notes[:i], notes[i+1:]...)

				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}
