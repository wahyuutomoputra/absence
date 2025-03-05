package repository

import (
	"absence/internal/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Create(ctx context.Context, attendance *model.Attendance) error
	GetByID(ctx context.Context, id uint) (*model.Attendance, error)
	GetByUserIDAndDate(ctx context.Context, userID uint, date time.Time) (*model.Attendance, error)
	Update(ctx context.Context, attendance *model.Attendance) error
	Delete(ctx context.Context, id uint) error
	GetUserAttendances(ctx context.Context, userID uint, startDate, endDate time.Time) ([]model.Attendance, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

func (r *attendanceRepository) Create(ctx context.Context, attendance *model.Attendance) error {
	return r.db.WithContext(ctx).Create(attendance).Error
}

func (r *attendanceRepository) GetByID(ctx context.Context, id uint) (*model.Attendance, error) {
	var attendance model.Attendance
	err := r.db.WithContext(ctx).First(&attendance, id).Error
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *attendanceRepository) GetByUserIDAndDate(ctx context.Context, userID uint, date time.Time) (*model.Attendance, error) {
	var attendance model.Attendance
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.db.WithContext(ctx).
		Where("user_id = ? AND check_in >= ? AND check_in < ?", userID, startOfDay, endOfDay).
		First(&attendance).Error
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *attendanceRepository) Update(ctx context.Context, attendance *model.Attendance) error {
	return r.db.WithContext(ctx).Save(attendance).Error
}

func (r *attendanceRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Attendance{}, id).Error
}

func (r *attendanceRepository) GetUserAttendances(ctx context.Context, userID uint, startDate, endDate time.Time) ([]model.Attendance, error) {
	var attendances []model.Attendance
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND check_in >= ? AND check_in <= ?", userID, startDate, endDate).
		Find(&attendances).Error
	return attendances, err
}
