package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JollyGrin/postgres-attendance/internal/api"
	"github.com/JollyGrin/postgres-attendance/internal/db"
	"github.com/JollyGrin/postgres-attendance/internal/model"
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

func (h *AttendanceHandler) RecordAttendance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		api.SendResponse(w, false, nil, "Method not allowed", "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Address        string `json:"address"`
		Location       string `json:"location"`
		Metaverse      string `json:"metaverse"`
		EntranceStatus string `json:"entrance_status"`
	}

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.SendResponse(w, false, nil, "Bad Request", "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate address (optional, but recommended)
	if req.Address == "" {
		api.SendResponse(w, false, nil, "Bad Request", "address is required", http.StatusBadRequest)
		return
	}

	if req.Location == "" {
		api.SendResponse(w, false, nil, "Bad Request", "location is required", http.StatusBadRequest)
		return
	}

	if req.EntranceStatus != "ENTER" && req.EntranceStatus != "EXIT" {
		api.SendResponse(w, false, nil, "Bad Request", "entrance_status must be 'ENTER' or 'EXIT'", http.StatusBadRequest)
		return
	}

	// Record the attendance
	err := h.db.RecordAttendance(r.Context(), req.Address, req.Location, model.MetaverseType(req.Metaverse), model.EntranceStatusType(req.EntranceStatus))
	if err != nil {
		errorMsg, details, statusCode := api.HandleDBError(fmt.Errorf("record attendance: %w", err))
		api.SendResponse(w, false, nil, errorMsg, details, statusCode)
		return
	}

	// Successful response
	api.SendResponse(w, true, nil, "", "Attendance recorded successfully", http.StatusCreated)
}
