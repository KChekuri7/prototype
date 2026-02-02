package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"movie-booking-api/src/models"
	"movie-booking-api/src/store"
)

type Handler struct {
	Store *store.Store
}

func NewHandler(s *store.Store) *Handler {
	return &Handler{Store: s}
}

// Helper to write JSON response
func (h *Handler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// Helper to write Error response
func (h *Handler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"error": message})
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies := h.Store.GetMovies()
	h.respondJSON(w, http.StatusOK, movies)
}

func (h *Handler) GetShows(w http.ResponseWriter, r *http.Request) {
	// Extract movieId from URL: /movies/{movieId}/shows
	// Since we are using standard http.ServeMux, we might need manual parsing if not using Go 1.22 path matching
	// Assuming Go 1.22+ "GET /movies/{id}/shows"
	
	movieId := r.PathValue("id") // Go 1.22 feature
	if movieId == "" {
		// Fallback for older Go versions if path value not available (handled by router config in main)
		// but let's assume we configure main.go correctly.
		// If manual parsing needed:
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) >= 4 && parts[1] == "movies" && parts[3] == "shows" {
			movieId = parts[2]
		}
	}

	shows := h.Store.GetShowsByMovie(movieId)
	h.respondJSON(w, http.StatusOK, shows)
}

func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	booking, err := h.Store.CreateBooking(req)
	if err != nil {
		if err == store.ErrShowNotFound {
			h.respondError(w, http.StatusNotFound, err.Error())
		} else if err == store.ErrSeatAlreadyBooked {
			h.respondError(w, http.StatusConflict, err.Error())
		} else {
			h.respondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	h.respondJSON(w, http.StatusCreated, booking)
}

func (h *Handler) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings := h.Store.GetBookings()
	h.respondJSON(w, http.StatusOK, bookings)
}
