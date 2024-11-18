package handler

import (
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
		errorMsg, details, statusCode := api.HandleDBError(err)
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
