package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "0.0.1"

type config struct {
	port int
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	// Try to read environment variable for port (given by railway). Otherwise use default
	port := os.Getenv("PORT")
	intPort, err := strconv.Atoi(port)
	if err != nil || intPort <= 0 {
		intPort = 4000
	}

	// Set the port to run the API on
	cfg.port = intPort

	// Create the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create the application
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Create the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  45 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("server started", "addr", srv.Addr)

	// Start the server
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("server error", "error", err)
		os.Exit(1)
	}
}

// Placeholder for the routes method
func (app *application) routes() http.Handler {
	// Define your routes here
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	return mux
}
