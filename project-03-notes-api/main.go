package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Notes API! | Hilal")
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the notes endpoint!")

	fmt.Fprintln(w, "HTTP Method:", r.Method)
	fmt.Fprintln(w, "URL Path:", r.URL.Path)
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
