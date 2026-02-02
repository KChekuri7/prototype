# Movie Ticket Booking API

A thread-safe REST API for booking movie tickets, built with Golang.

## Setup & Run

Prerequisites: Go 1.22+

1.  **Initialize Module** (if not already done)
    ```bash
    go mod tidy
    ```

2.  **Run Server**
    ```bash
    go run main.go
    ```
    Server starts on `http://localhost:8080`.

3.  **Run Tests**
    ```bash
    go test -v
    ```

## API Endpoints

### 1. Health Check
**Request:**
`GET /health`

**Response:**
```json
{
  "status": "ok"
}
```

### 2. List Movies
**Request:**
`GET /movies`

**Response:**
```json
[
  {
    "id": "m1",
    "title": "The Godfather",
    "duration": 175,
    "rating": "R"
  },
  ...
]
```

### 3. List Shows
**Request:**
`GET /movies/m1/shows`

**Response:**
```json
[
  {
    "id": "s1",
    "movieId": "m1",
    "time": "2026-02-05T18:00:00Z",
    "totalSeats": 20,
    "bookedSeats": ["A1", "A2"]
  }
]
```

### 4. Create Booking
**Request:**
`POST /bookings`
```json
{
  "showId": "s1",
  "seats": ["B1", "B2"]
}
```

**Response (Success - 201 Created):**
```json
{
  "id": "20260205120000",
  "showId": "s1",
  "seats": ["B1", "B2"],
  "createdAt": "2026-02-05T12:00:00Z"
}
```

**Response (Failure - 409 Conflict):**
```json
{
  "error": "one or more seats are already booked"
}
```

## Features
- **In-Memory Storage**: Data is stored in memory and reset on restart.
- **Thread Safety**: Uses `sync.RWMutex` to handle concurrent bookings safely.
- **Validation**: Prevents double booking and invalid show IDs.
