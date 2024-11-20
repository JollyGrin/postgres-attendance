package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JollyGrin/postgres-attendance/internal/db"
	"github.com/JollyGrin/postgres-attendance/internal/handler"
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
	http.HandleFunc("/api/attendance/by", attendanceHandler.GetAttendanceByAddress)
	http.HandleFunc("/api/attendance/date", attendanceHandler.GetUniqueAddressesByDay)

	// Start server
	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
