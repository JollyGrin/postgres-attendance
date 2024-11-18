package main

import (
	"fmt"
	"github.com/JollyGrin/postgres-attendance/internal/db"
	"github.com/JollyGrin/postgres-attendance/internal/handler"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
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

	// Initialize database connection
	database, err := db.NewDB(dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer database.Close()

	// Initialize handler with database pool
	attendanceHandler := handler.NewAttendanceHandler(database)

	// Define routes
	http.HandleFunc("/api/attendance/today", attendanceHandler.GetTodayAttendance)
	// http.HandleFunc("/api/attendance", handler.createAttendance)

	// Start server
	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func (h *AttendanceHandler) createAttendance(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		sendResponse(w, false, nil, "Method not allowed", "Only POST method is allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var attendance Attendance
// 	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
// 		sendResponse(w, false, nil, "Invalid request body", "The provided JSON is malformed or invalid", http.StatusBadRequest)
// 		return
// 	}

// 	// Validate required fields
// 	if attendance.Address == "" {
// 		sendResponse(w, false, nil, "Missing required field", "Address field is required", http.StatusBadRequest)
// 		return
// 	}

// 	// Insert the new attendance record
// 	var id string
// 	err := h.pool.QueryRow(context.Background(),
// 		"INSERT INTO attendance (address) VALUES ($1) RETURNING id",
// 		attendance.Address, attendance.Created_At).Scan(&id)

// 	if err != nil {
// 		errorMsg, details, statusCode := handleDBError(err)
// 		sendResponse(w, false, nil, errorMsg, details, statusCode)
// 		return
// 	}

// 	attendance.ID = id
// 	sendResponse(w, true, attendance, "", "", http.StatusCreated)
// }

// sendResponse is a helper function to send JSON responses
