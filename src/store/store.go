package store

import (
	"errors"
	"sync"
	"time"

	"movie-booking-api/src/models"
)

var (
	ErrShowNotFound      = errors.New("show not found")
	ErrSeatAlreadyBooked = errors.New("one or more seats are already booked")
	ErrInvalidSeat       = errors.New("invalid seat")
)

// Store holds the in-memory data
type Store struct {
	movies   map[string]models.Movie
	shows    map[string]*models.Show // Pointer to allow modification
	bookings map[string]models.Booking
	mu       sync.RWMutex // Protects all maps (global lock for simplicity, or per-show lock preferable for scale)
}

// NewStore initializes the store with some seed data
func NewStore() *Store {
	s := &Store{
		movies:   make(map[string]models.Movie),
		shows:    make(map[string]*models.Show),
		bookings: make(map[string]models.Booking),
	}
	s.seed()
	return s
}

func (s *Store) seed() {
	// Seed Movies
	m1 := models.Movie{ID: "m1", Title: "The Godfather", Duration: 175, Rating: "R"}
	m2 := models.Movie{ID: "m2", Title: "The Dark Knight", Duration: 152, Rating: "PG-13"}
	s.movies[m1.ID] = m1
	s.movies[m2.ID] = m2

	// Seed Shows
	show1 := &models.Show{
		ID:          "s1",
		MovieID:     "m1",
		Time:        time.Date(2026, 2, 5, 18, 0, 0, 0, time.UTC),
		TotalSeats:  20,
		BookedSeats: []string{"A1", "A2"},
	}
	show2 := &models.Show{
		ID:          "s2",
		MovieID:     "m2",
		Time:        time.Date(2026, 2, 5, 21, 0, 0, 0, time.UTC),
		TotalSeats:  20,
		BookedSeats: []string{},
	}
	s.shows[show1.ID] = show1
	s.shows[show2.ID] = show2
}

// GetMovies returns all movies
func (s *Store) GetMovies() []models.Movie {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var movies []models.Movie
	for _, m := range s.movies {
		movies = append(movies, m)
	}
	return movies
}

// GetShowsByMovie returns all shows for a given movie ID
func (s *Store) GetShowsByMovie(movieID string) []models.Show {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var shows []models.Show
	for _, show := range s.shows {
		if show.MovieID == movieID {
			shows = append(shows, *show)
		}
	}
	return shows
}

// GetBookings returns all bookings
func (s *Store) GetBookings() []models.Booking {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var bookings []models.Booking
	for _, b := range s.bookings {
		bookings = append(bookings, b)
	}
	return bookings
}

// CreateBooking attempts to book seats for a show
func (s *Store) CreateBooking(req models.CreateBookingRequest) (*models.Booking, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	show, exists := s.shows[req.ShowID]
	if !exists {
		return nil, ErrShowNotFound
	}

	// Check availability
	// Logic: Iterate requested seats, check if exists in BookedSeats
	for _, requestedSeat := range req.Seats {
		for _, bookedSeat := range show.BookedSeats {
			if requestedSeat == bookedSeat {
				return nil, ErrSeatAlreadyBooked
			}
		}
	}

	// Book the seats
	show.BookedSeats = append(show.BookedSeats, req.Seats...)

	// Create request
	booking := models.Booking{
		ID:        generateID(),
		ShowID:    req.ShowID,
		Seats:     req.Seats,
		CreatedAt: time.Now(),
	}
	s.bookings[booking.ID] = booking

	return &booking, nil
}

func generateID() string {
	// Simple ID generation
	return time.Now().Format("20060102150405")
}
