package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	logger *slog.Logger
}

func main() {
	// run on the port given through -addr flag i.e. -addr=":9999"
	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()

	// use structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// init instance of app struct that contains dependencies
	app := &Application{
		logger: logger,
	}

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

	// logs and errors
	logger.Info("starting server...", "addr", *addr)

	// listen on localhost:addr
	err := http.ListenAndServe(*addr, mux)

	// log the error returned and terminate app
	logger.Error(err.Error())
	os.Exit(1)
}
