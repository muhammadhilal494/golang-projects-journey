package main

import "strings"

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func validateNote(note *Note) bool {

	note.Title = strings.TrimSpace(note.Title)
	note.Content = strings.TrimSpace(note.Content)

	if note.Title == "" || note.Content == "" {
		return false
	}

	return true
}
