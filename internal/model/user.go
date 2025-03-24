package model

import (
	"time"
)

// User represents the user model
// @Description User model
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null;size:50" json:"username"`
	Password  string    `gorm:"not null;size:255" json:"-"`
	FullName  string    `gorm:"not null;size:100" json:"full_name"`
	Email     string    `gorm:"unique;not null;size:100" json:"email"`
	Role      string    `gorm:"not null;size:20;check:role IN ('admin', 'employee')" json:"role"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}
