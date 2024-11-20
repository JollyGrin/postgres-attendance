package handler

import (
	"fmt"
	"net/http"

	"github.com/JollyGrin/postgres-attendance/internal/api"
	"github.com/JollyGrin/postgres-attendance/internal/db"
)

type AttendanceHandler struct {
	db *db.DB
}

func NewAttendanceHandler(db *db.DB) *AttendanceHandler {
	return &AttendanceHandler{db: db}
}

func (h *AttendanceHandler) GetTodayAttendance(w http.ResponseWriter, r *http.Request) {
	// Method check
	if r.Method != http.MethodGet {
		api.SendResponse(w, false, nil, "Method not allowed", "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get records from database
	records, err := h.db.GetTodayAttendance(r.Context())
	if err != nil {
		errorMsg, details, statusCode := api.HandleDBError(fmt.Errorf("get attendance: %w", err))
		api.SendResponse(w, false, nil, errorMsg, details, statusCode)
		return
	}

	// Handle empty results
	if len(records) == 0 {
		api.SendResponse(w, true, records, "", "No attendance records found for today", http.StatusOK)
		return
	}

	// Success response
	api.SendResponse(w, true, records, "", "", http.StatusOK)
}

func (h *AttendanceHandler) GetAttendanceByAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.SendResponse(w, false, nil, "Method not allowed", "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the address from query parameters
	address := r.URL.Query().Get("address")
	if address == "" {
		api.SendResponse(w, false, nil, "Bad Request", "Address parameter is required", http.StatusBadRequest)
		return
	}

	records, err := h.db.GetRecordsByAddress(r.Context(), address)

	if err != nil {
		errorMsg, details, statusCode := api.HandleDBError(fmt.Errorf("get attendance by record: %w", err))
		api.SendResponse(w, false, nil, errorMsg, details, statusCode)
		return
	}

	if len(records) == 0 {
		api.SendResponse(w, true, records, "", "No attendance records found for address", http.StatusOK)
		return
	}

	api.SendResponse(w, true, records, "", "", http.StatusOK)
}

func (h *AttendanceHandler) GetUniqueAddressesByDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.SendResponse(w, false, nil, "Method not allowed", "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	day := r.URL.Query().Get("day")
	if day == "" {
		api.SendResponse(w, false, nil, "Bad Request", "Day parameter is required", http.StatusBadRequest)
		return
	}

	uniqueCount, uniqueAddresses, err := h.db.GetUniqueAddressesByDay(r.Context(), day)
	if err != nil {
		errorMsg, details, statusCode := api.HandleDBError(fmt.Errorf("get attendance by day: %w", err))
		api.SendResponse(w, false, nil, errorMsg, details, statusCode)
		return
	}

	if uniqueCount == 0 {
		api.SendResponse(w, true, uniqueAddresses, "", "No attendance records found for day", http.StatusOK)
		return
	}

	api.SendResponse(w, true, uniqueAddresses, "", "", http.StatusOK)
}
