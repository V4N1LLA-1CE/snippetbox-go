package main

import (
	"database/sql"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Application struct {
	logger *slog.Logger
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("cannot load .env variables: %v\n", err)
	}
}

func main() {
	// run on the port given through -addr flag i.e. -addr=":9999"
	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()

	// use structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := openPostgresDB(os.Getenv("POSTGRES_URL"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// init instance of app struct that contains dependencies
	app := &Application{
		logger: logger,
	}

	// logs and errors
	logger.Info("starting server...", "addr", *addr)

	// listen on localhost:addr
	err = http.ListenAndServe(*addr, app.router())

	// log the error returned and terminate app
	logger.Error(err.Error())
	os.Exit(1)
}

func openPostgresDB(dsn string) (*sql.DB, error) {
	// initialise db pool
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// establish connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
