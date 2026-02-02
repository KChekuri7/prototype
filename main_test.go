package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"movie-booking-api/src/controllers"
	"movie-booking-api/src/models"
	"movie-booking-api/src/store"
)

func setupTestServer() *http.ServeMux {
	s := store.NewStore()
	h := controllers.NewHandler(s)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("GET /movies", h.GetMovies)
	mux.HandleFunc("GET /movies/{id}/shows", h.GetShows)
	mux.HandleFunc("POST /bookings", h.CreateBooking)
	return mux
}

func TestHealthCheck(t *testing.T) {
	mux := setupTestServer()
	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("health check returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetMovies(t *testing.T) {
	mux := setupTestServer()
	req, _ := http.NewRequest("GET", "/movies", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200, got %v", status)
	}

	var movies []models.Movie
	json.Unmarshal(rr.Body.Bytes(), &movies)
	if len(movies) == 0 {
		t.Errorf("expected movies, got empty list")
	}
}

func TestBookingConcurrency(t *testing.T) {
	// Setup fresh server/store
	s := store.NewStore()
	h := controllers.NewHandler(s)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /bookings", h.CreateBooking)

	// We'll target show "s1" which has "A1", "A2" booked initially.
	// Total seats 20.
	// We will try to confirm that double booking "B1" is prevented.

	targetSeat := "B1"
	showID := "s1"

	var wg sync.WaitGroup
	concurrentRequests := 10
	successCount := 0
	failCount := 0
	var mu sync.Mutex

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			payload := models.CreateBookingRequest{
				ShowID: showID,
				Seats:  []string{targetSeat},
			}
			body, _ := json.Marshal(payload)
			req, _ := http.NewRequest("POST", "/bookings", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			mu.Lock()
			if rr.Code == http.StatusCreated {
				successCount++
			} else if rr.Code == http.StatusConflict {
				failCount++
			}
			mu.Unlock()
		}()
	}

	wg.Wait()

	if successCount != 1 {
		t.Errorf("Expected exactly 1 successful booking, got %d", successCount)
	}
	if failCount != concurrentRequests-1 {
		t.Errorf("Expected %d failed bookings, got %d", concurrentRequests-1, failCount)
	}
}
