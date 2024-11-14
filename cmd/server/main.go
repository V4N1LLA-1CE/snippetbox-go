package main

import (
	"log"
	"net/http"
)

func main() {
	// initialise router
	mux := http.NewServeMux()

	// initialise static fileserver to serve files from "/ui/static/" directory
	// path given to http.Dir is relative to project root
	fs := http.FileServer(http.Dir(("./ui/static/")))

	// Map URL starting with "/static/" to the file server
	// i.e. "/static/" is mapped to "./ui/static/" here
	// Strip "/static" from full path URL before looking for file
	mux.Handle("GET /static/", http.StripPrefix("/static", fs))

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
