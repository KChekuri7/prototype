package main

import (
	"log"
	"net/http"
	"time"

	"movie-booking-api/src/controllers"
	"movie-booking-api/src/store"
)

func main() {
	// Initialize Store
	s := store.NewStore()

	// Initialize Handlers
	h := controllers.NewHandler(s)

	// Router
	mux := http.NewServeMux()

	// Register Routes
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("GET /movies", h.GetMovies)
	mux.HandleFunc("GET /movies/{id}/shows", h.GetShows)
	mux.HandleFunc("POST /bookings", h.CreateBooking)
	mux.HandleFunc("GET /bookings", h.GetBookings)

	// Server config
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
