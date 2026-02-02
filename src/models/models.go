package models

import "time"

// Movie represents a movie in the cinema
type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"` // in minutes
	Rating   string `json:"rating"`
}

// Show represents a specific screening of a movie
type Show struct {
	ID          string    `json:"id"`
	MovieID     string    `json:"movieId"`
	Time        time.Time `json:"time"`
	TotalSeats  int       `json:"totalSeats"`
	BookedSeats []string  `json:"bookedSeats"`
}

// Booking represents a user's ticket booking
type Booking struct {
	ID        string    `json:"id"`
	ShowID    string    `json:"showId"`
	Seats     []string  `json:"seats"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateBookingRequest is the payload for booking seats
type CreateBookingRequest struct {
	ShowID string   `json:"showId"`
	Seats  []string `json:"seats"`
}
