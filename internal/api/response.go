package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/lib/pq"
)

type Response struct {
	Success bool        `json:"success"` // TODO: Remove  this, use status codes to indicate success/failure
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`         // TODO: Rename to "message", if failure, then it is the error message
	Details string      `json:"error_details,omitempty"` // Will be the golang error on fails
}

// sendResponse is a helper function to send JSON responses
// TODO: rename to "write" instead of send
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
	// var errorMsg, details string
	// var statusCode int

	if errors.Is(err, sql.ErrNoRows) {
		return "Record not found", err.Error(), http.StatusNotFound
	}

	var pqErr *pq.Error
	if !errors.As(err, &pqErr) {
		return "Database error occurred", err.Error(), http.StatusInternalServerError
	}

	switch pqErr.Code.Name() {
	case "unique_violation":
		return "Duplicate record", err.Error(), http.StatusConflict

	}
	return fmt.Sprintf("Database error %q", pqErr.Code.Name()), err.Error(), http.StatusBadRequest

}
