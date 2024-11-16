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

	// logs and errors
	logger.Info("starting server...", "addr", *addr)

	// listen on localhost:addr
	err := http.ListenAndServe(*addr, app.router())

	// log the error returned and terminate app
	logger.Error(err.Error())
	os.Exit(1)
}
