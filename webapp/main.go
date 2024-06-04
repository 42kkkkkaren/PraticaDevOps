package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static assets like CSS and JavaScript
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Serve HTML files
	http.Handle("/", http.FileServer(http.Dir("views")))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
