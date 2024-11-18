package api

import (
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Details string      `json:"error_details,omitempty"`
}

// sendResponse is a helper function to send JSON responses
func SendResponse(w http.ResponseWriter, success bool, data interface{}, errorMsg, details string, statusCode int) {
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

func HandleDBError(err error) (string, string, int) {
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
