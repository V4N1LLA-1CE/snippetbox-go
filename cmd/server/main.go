package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// run on the port given through -addr flag i.e. -addr=":9999"
	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()

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
	log.Printf("starting server on %v\n", *addr)

	// listen on localhost:4000
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
