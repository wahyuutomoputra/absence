package handler

import (
	"absence/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	attendanceService service.AttendanceService
}

func NewAttendanceHandler(attendanceService service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{attendanceService: attendanceService}
}

// CheckIn godoc
// @Summary Record attendance check-in
// @Description Record user's check-in time and location
// @Tags attendance
// @Accept json
// @Produce json
// @Param location body object{location=string} true "Check-in location"
// @Success 200 {object} map[string]string "Check-in successful"
// @Failure 400 {object} map[string]string "Invalid input or already checked in"
// @Security BearerAuth
// @Router /attendance/check-in [post]
func (h *AttendanceHandler) CheckIn(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request struct {
		Location string `json:"location" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.attendanceService.CheckIn(c.Request.Context(), uint(userID), request.Location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Check-in successful"})
}

// CheckOut godoc
// @Summary Record attendance check-out
// @Description Record user's check-out time and location
// @Tags attendance
// @Accept json
// @Produce json
// @Param location body object{location=string} true "Check-out location"
// @Success 200 {object} map[string]string "Check-out successful"
// @Failure 400 {object} map[string]string "Invalid input or no check-in record"
// @Security BearerAuth
// @Router /attendance/check-out [post]
func (h *AttendanceHandler) CheckOut(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request struct {
		Location string `json:"location" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.attendanceService.CheckOut(c.Request.Context(), uint(userID), request.Location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Check-out successful"})
}

// GetAttendance godoc
// @Summary Get attendance record
// @Description Get attendance record by ID
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} model.Attendance "Attendance record"
// @Failure 400 {object} map[string]string "Invalid attendance ID"
// @Failure 404 {object} map[string]string "Attendance record not found"
// @Security BearerAuth
// @Router /attendance/{id} [get]
func (h *AttendanceHandler) GetAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendance ID"})
		return
	}

	attendance, err := h.attendanceService.GetAttendanceByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendance record not found"})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// GetUserAttendances godoc
// @Summary Get user's attendance records
// @Description Get user's attendance records within a date range
// @Tags attendance
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {array} model.Attendance "List of attendance records"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Server error"
// @Security BearerAuth
// @Router /users/{user_id}/attendance [get]
func (h *AttendanceHandler) GetUserAttendances(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request struct {
		StartDate string `form:"start_date" binding:"required"`
		EndDate   string `form:"end_date" binding:"required"`
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	attendances, err := h.attendanceService.GetUserAttendances(c.Request.Context(), uint(userID), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendances)
}
