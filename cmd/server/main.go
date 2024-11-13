package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a dorm for creating a new snippet..."))
}

func main() {
	// initialise new router
	// register homeHandler
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", homeHandler) // Restrict route to only match on "/"
	mux.HandleFunc("/snippetbox/view", viewHandler)
	mux.HandleFunc("/snippetbox/create", createHandler)

	// log
	log.Println("Starting server on :4000")

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
