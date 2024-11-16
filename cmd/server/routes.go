package main

import "net/http"

func (app *Application) router() *http.ServeMux {
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
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
