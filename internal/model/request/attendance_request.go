package request

// CheckInRequest represents the request body for check-in
type CheckInRequest struct {
	// Location is optional and can be used to store check-in location
	Location string `json:"location" example:"Jakarta"`
	// Notes is optional and can be used to store additional information
	Notes string `json:"notes" example:"Working from home"`
}

// CheckOutRequest represents the request body for check-out
type CheckOutRequest struct {
	// Location is optional and can be used to store check-out location
	Location string `json:"location" example:"Jakarta"`
	// Notes is optional and can be used to store additional information
	Notes string `json:"notes" example:"Finished work for today"`
}

// GetAttendanceRequest represents the request parameters for getting attendance
type GetAttendanceRequest struct {
	// Date is optional and can be used to filter attendance by date (format: YYYY-MM-DD)
	Date string `json:"date" example:"2024-03-20"`
	// Month is optional and can be used to filter attendance by month (format: YYYY-MM)
	Month string `json:"month" example:"2024-03"`
	// Year is optional and can be used to filter attendance by year (format: YYYY)
	Year string `json:"year" example:"2024"`
}
