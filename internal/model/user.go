package model

import (
	"time"
)

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

// RegisterRequest represents the user registration request
// @Description User registration request
type RegisterRequest struct {
	Username string `json:"username" example:"john_doe" binding:"required"`
	Password string `json:"password" example:"secure123" binding:"required"`
	FullName string `json:"full_name" example:"John Doe" binding:"required"`
	Email    string `json:"email" example:"john@example.com" binding:"required,email"`
	Role     string `json:"role" example:"employee" binding:"required,oneof=admin employee"`
}

// LoginRequest represents the login credentials
// @Description Login request
type LoginRequest struct {
	Username string `json:"username" example:"john_doe" binding:"required"`
	Password string `json:"password" example:"secure123" binding:"required"`
}

// UpdateUserRequest represents the user update request
// @Description User update request
type UpdateUserRequest struct {
	Username string `json:"username" example:"john_doe" binding:"required"`
	FullName string `json:"full_name" example:"John Doe Updated" binding:"required"`
	Email    string `json:"email" example:"john.updated@example.com" binding:"required,email"`
	Role     string `json:"role" example:"employee" binding:"required,oneof=admin employee"`
}
