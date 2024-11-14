package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %v\n", id)
	w.Write([]byte(msg))
}

func createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func createSnippetPostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	// initialise new router
	// register homeHandler
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", homeHandler) // Restrict route to only match on "/"
	mux.HandleFunc("GET /snippetbox/view/{id}", viewHandler)
	mux.HandleFunc("GET /snippetbox/create", createSnippetHandler)
	mux.HandleFunc("POST /snippetbox/create", createSnippetPostHandler)

	// log
	log.Println("Starting server on :4000")

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
