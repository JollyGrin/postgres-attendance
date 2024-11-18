package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Attendance represents the structure of our attendance record
type Attendance struct {
	ID         string    `json:"id"`
	Address    string    `json:"address"`
	Created_At time.Time `json:"created_at"`
}

// AttendanceHandler handles attendance-related requests
type AttendanceHandler struct {
	pool *pgxpool.Pool
}

// Response represents the structure of our API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Details string      `json:"error_details,omitempty"` // Added field for detailed error messages
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	// Build the database connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Connect to PostgreSQL
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	fmt.Println("Connected to PostgreSQL successfully!")

	// Initialize handler with database pool
	handler := &AttendanceHandler{pool: pool}

	// Define routes
	http.HandleFunc("/api/attendance/today", handler.getTodayAttendance)
	http.HandleFunc("/api/attendance", handler.createAttendance)

	// Start server
	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleDBError processes database errors and returns appropriate error messages and status codes
func handleDBError(err error) (string, string, int) {
	var errorMsg, details string
	var statusCode int

	switch {
	case err == pgx.ErrNoRows:
		errorMsg = "No records found"
		details = "The query returned no results"
		statusCode = http.StatusNotFound
	case err.Error() == "ERROR: duplicate key value violates unique constraint" ||
		err.Error() == "ERROR: unique_violation":
		errorMsg = "Duplicate record"
		details = "A record with this key already exists"
		statusCode = http.StatusConflict
	case err.Error() == "ERROR: foreign key violation" ||
		err.Error() == "ERROR: foreign_key_violation":
		errorMsg = "Invalid reference"
		details = "The record references a non-existent related record"
		statusCode = http.StatusBadRequest
	case err.Error() == "ERROR: null value in column violates not-null constraint":
		errorMsg = "Missing required field"
		details = "A required field was not provided"
		statusCode = http.StatusBadRequest
	default:
		errorMsg = "Database error occurred"
		details = err.Error()
		statusCode = http.StatusInternalServerError
	}

	return errorMsg, details, statusCode
}

// getTodayAttendance handles GET requests to fetch today's attendance records
func (h *AttendanceHandler) getTodayAttendance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendResponse(w, false, nil, "Method not allowed", "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get today's date in YYYY-MM-DD format
	today := time.Now().Format("2006-01-02")

	// Query today's attendance records
	rows, err := h.pool.Query(context.Background(),
		"SELECT id, address, created_at FROM attendance WHERE DATE(created_at) = $1",
		today)
	if err != nil {
		errorMsg, details, statusCode := handleDBError(err)
		sendResponse(w, false, nil, errorMsg, details, statusCode)
		return
	}
	defer rows.Close()

	var records []Attendance
	for rows.Next() {
		var record Attendance
		err := rows.Scan(&record.ID, &record.Address, &record.Created_At)
		if err != nil {
			errorMsg, details, statusCode := handleDBError(err)
			sendResponse(w, false, nil, errorMsg, details, statusCode)
			return
		}
		records = append(records, record)
	}

	if err = rows.Err(); err != nil {
		errorMsg, details, statusCode := handleDBError(err)
		sendResponse(w, false, nil, errorMsg, details, statusCode)
		return
	}

	// If no records found, return a specific message
	if len(records) == 0 {
		sendResponse(w, true, records, "", "No attendance records found for today", http.StatusOK)
		return
	}

	sendResponse(w, true, records, "", "", http.StatusOK)
}

// createAttendance handles POST requests to create new attendance records
func (h *AttendanceHandler) createAttendance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendResponse(w, false, nil, "Method not allowed", "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var attendance Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		sendResponse(w, false, nil, "Invalid request body", "The provided JSON is malformed or invalid", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if attendance.Address == "" {
		sendResponse(w, false, nil, "Missing required field", "Address field is required", http.StatusBadRequest)
		return
	}

	// Insert the new attendance record
	var id string
	err := h.pool.QueryRow(context.Background(),
		"INSERT INTO attendance (address) VALUES ($1) RETURNING id",
		attendance.Address, attendance.Created_At).Scan(&id)

	if err != nil {
		errorMsg, details, statusCode := handleDBError(err)
		sendResponse(w, false, nil, errorMsg, details, statusCode)
		return
	}

	attendance.ID = id
	sendResponse(w, true, attendance, "", "", http.StatusCreated)
}

// sendResponse is a helper function to send JSON responses
func sendResponse(w http.ResponseWriter, success bool, data interface{}, errorMsg, details string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: success,
		Data:    data,
		Error:   errorMsg,
		Details: details,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
