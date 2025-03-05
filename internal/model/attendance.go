package model

import (
	"time"
)

type Attendance struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	CheckIn     time.Time `json:"check_in"`
	CheckOut    time.Time `json:"check_out"`
	Status      string    `gorm:"size:20;check:status IN ('present', 'late', 'absent', 'leave')" json:"status"`
	LocationIn  string    `gorm:"type:point" json:"location_in"`
	LocationOut string    `gorm:"type:point" json:"location_out"`
	Notes       string    `gorm:"type:text" json:"notes"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}
