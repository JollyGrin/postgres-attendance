package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

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

	fmt.Println("Connected to PostgreSQL successfully!")

	defer pool.Close()

	fmt.Println("Connected to PostgreSQL successfully!")

	// Query the attendance table
	rows, err := pool.Query(context.Background(), "SELECT * FROM attendance")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	fmt.Println("Attendance Records:")
	for rows.Next() {
		var id int
		var name string
		var date string
		err := rows.Scan(&id, &name, &date)
		if err != nil {
			log.Fatalf("Error scanning row: %v\n", err)
		}
		fmt.Printf("ID: %d, Name: %s, Date: %s\n", id, name, date)
	}

	if rows.Err() != nil {
		log.Fatalf("Error during rows iteration: %v\n", rows.Err())
	}
}
