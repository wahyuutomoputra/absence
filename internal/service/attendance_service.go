package service

import (
	"absence/internal/model"
	"absence/internal/repository"
	"context"
	"errors"
	"time"
)

type AttendanceService interface {
	CheckIn(ctx context.Context, userID uint, location string) error
	CheckOut(ctx context.Context, userID uint, location string) error
	GetAttendanceByID(ctx context.Context, id uint) (*model.Attendance, error)
	GetUserAttendances(ctx context.Context, userID uint, startDate, endDate time.Time) ([]model.Attendance, error)
}

type attendanceService struct {
	attendanceRepo repository.AttendanceRepository
}

func NewAttendanceService(attendanceRepo repository.AttendanceRepository) AttendanceService {
	return &attendanceService{attendanceRepo: attendanceRepo}
}

func (s *attendanceService) CheckIn(ctx context.Context, userID uint, location string) error {
	now := time.Now()

	// Check if already checked in today
	existing, _ := s.attendanceRepo.GetByUserIDAndDate(ctx, userID, now)
	if existing != nil {
		return errors.New("already checked in today")
	}

	attendance := &model.Attendance{
		UserID:   userID,
		CheckIn:  now,
		Location: location,
	}

	return s.attendanceRepo.Create(ctx, attendance)
}

func (s *attendanceService) CheckOut(ctx context.Context, userID uint, location string) error {
	now := time.Now()

	attendance, err := s.attendanceRepo.GetByUserIDAndDate(ctx, userID, now)
	if err != nil {
		return errors.New("no check-in record found for today")
	}

	attendance.CheckOut = now
	attendance.Location = location

	return s.attendanceRepo.Update(ctx, attendance)
}

func (s *attendanceService) GetAttendanceByID(ctx context.Context, id uint) (*model.Attendance, error) {
	return s.attendanceRepo.GetByID(ctx, id)
}

func (s *attendanceService) GetUserAttendances(ctx context.Context, userID uint, startDate, endDate time.Time) ([]model.Attendance, error) {
	return s.attendanceRepo.GetUserAttendances(ctx, userID, startDate, endDate)
}
