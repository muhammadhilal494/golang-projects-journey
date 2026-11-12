package main

import (
	"fmt"
	"net/http"

	"github.com/muhammadhilal494/golang-projects-journey/project-03-notes-api/internal/note"
)

var notes []note.Note
var nextNoteID = 1

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
