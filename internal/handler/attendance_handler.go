package handler

import (
	"absence/internal/model/request"
	"absence/internal/service"
	"absence/pkg/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	attendanceService service.AttendanceService
}

func NewAttendanceHandler(attendanceService service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
	}
}

// CheckIn godoc
// @Summary Check-in attendance
// @Description Record user check-in for the day
// @Tags attendance
// @Accept json
// @Produce json
// @Param request body request.CheckInRequest true "Check-in details"
// @Success 200 {object} response.Response{data=model.Attendance} "Check-in successful"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 409 {object} response.Response "Already checked in"
// @Security BearerAuth
// @Router /attendance/check-in [post]
func (h *AttendanceHandler) CheckIn(c *gin.Context) {
	var req request.CheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.attendanceService.CheckIn(c.Request.Context(), userID.(uint), req.Location); err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Check-in successful", nil)
}

// CheckOut godoc
// @Summary Check-out attendance
// @Description Record user check-out for the day
// @Tags attendance
// @Accept json
// @Produce json
// @Param request body request.CheckOutRequest true "Check-out details"
// @Success 200 {object} response.Response{data=model.Attendance} "Check-out successful"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 404 {object} response.Response "No check-in record found"
// @Security BearerAuth
// @Router /attendance/check-out [post]
func (h *AttendanceHandler) CheckOut(c *gin.Context) {
	var req request.CheckOutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.attendanceService.CheckOut(c.Request.Context(), userID.(uint), req.Location); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Check-out successful", nil)
}

// GetAttendance godoc
// @Summary Get attendance by ID
// @Description Get attendance details by ID
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} response.Response{data=model.Attendance} "Attendance details retrieved successfully"
// @Failure 400 {object} response.Response "Invalid attendance ID"
// @Failure 404 {object} response.Response "Attendance not found"
// @Security BearerAuth
// @Router /attendance/{id} [get]
func (h *AttendanceHandler) GetAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	attendance, err := h.attendanceService.GetAttendanceByID(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Attendance not found")
		return
	}

	response.Success(c, http.StatusOK, "Attendance details retrieved successfully", attendance)
}

// GetUserAttendances godoc
// @Summary Get user attendances
// @Description Get attendance history for a user
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param date query string false "Filter by date (YYYY-MM-DD)"
// @Param month query string false "Filter by month (YYYY-MM)"
// @Param year query string false "Filter by year (YYYY)"
// @Success 200 {object} response.Response{data=[]model.Attendance} "User attendances retrieved successfully"
// @Failure 400 {object} response.Response "Invalid user ID or date format"
// @Security BearerAuth
// @Router /users/{id}/attendance [get]
func (h *AttendanceHandler) GetUserAttendances(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var startDate, endDate time.Time
	now := time.Now()

	// Parse date parameters
	date := c.Query("date")
	month := c.Query("month")
	year := c.Query("year")

	if date != "" {
		startDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid date format")
			return
		}
		endDate = startDate.Add(24 * time.Hour)
	} else if month != "" {
		startDate, err = time.Parse("2006-01", month)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid month format")
			return
		}
		endDate = startDate.AddDate(0, 1, 0)
	} else if year != "" {
		startDate, err = time.Parse("2006", year)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid year format")
			return
		}
		endDate = startDate.AddDate(1, 0, 0)
	} else {
		// Default to current month
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		endDate = startDate.AddDate(0, 1, 0)
	}

	attendances, err := h.attendanceService.GetUserAttendances(c.Request.Context(), uint(userID), startDate, endDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User attendances retrieved successfully", attendances)
}
