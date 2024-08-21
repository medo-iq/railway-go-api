package main

import (
	"encoding/json"
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

type ResponseData struct {
	Message    string    `json:"message"`
	Data       []string  `json:"data"`
	FetchTime  float64   `json:"fetch_time"` // Time in milliseconds
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
		Handler:      app.routes(), // Ensure this method is defined correctly
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

// Define the routes method
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/data", app.dataHandler)
	return mux
}

func (app *application) dataHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Simulate data fetching
	data := []string{"item1", "item2", "item3"}

	// Measure fetch time
	fetchTime := time.Since(startTime).Seconds() * 1000 // Time in milliseconds

	response := ResponseData{
		Message:   "Data fetched successfully",
		Data:      data,
		FetchTime: fetchTime,
	}

	// Set response header and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.logger.Error("failed to encode response", "error", err)
	}
}
