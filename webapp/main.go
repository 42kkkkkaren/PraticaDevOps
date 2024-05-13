package main

import (
	"net/http"
)

func main() {
	// Serve static assets like CSS and JavaScript
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Serve HTML files
	http.Handle("/", http.FileServer(http.Dir("views")))

	http.ListenAndServe(":8080", nil)
}
