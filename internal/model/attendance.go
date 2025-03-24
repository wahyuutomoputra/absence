package model

import (
	"time"
)

// Attendance represents the attendance record in the system
type Attendance struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	CheckIn   time.Time `json:"check_in" gorm:"not null"`
	CheckOut  time.Time `json:"check_out"`
	Location  string    `json:"location"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Attendance
func (Attendance) TableName() string {
	return "attendances"
}
