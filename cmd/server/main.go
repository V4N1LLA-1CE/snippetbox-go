package main

import (
	"log"
	"net/http"
)

func main() {
	// initialise router
	mux := http.NewServeMux()

	// create endpoints
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// logs and errors
	log.Print("starting server on :4000")

	// listen on localhost:4000
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
