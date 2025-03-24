//go:build wireinject
// +build wireinject

package internal

import (
	"absence/internal/handler"
	"absence/internal/middleware"
	"absence/internal/repository"
	"absence/internal/service"
	"absence/pkg/database"
	"absence/pkg/jwt"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitializeDB initializes the database connection
func InitializeDB(config *database.Config) (*gorm.DB, error) {
	return database.NewDatabase(config)
}

// InitializeAPI initializes all components of the API
func InitializeAPI(db *gorm.DB, jwtManager *jwt.JWTManager) (*API, error) {
	wire.Build(
		repository.NewUserRepository,
		repository.NewAttendanceRepository,
		service.NewUserService,
		service.NewAttendanceService,
		handler.NewUserHandler,
		handler.NewAttendanceHandler,
		middleware.NewAuthMiddleware,
		wire.Struct(new(API), "*"),
	)
	return nil, nil
}

type API struct {
	UserHandler       *handler.UserHandler
	AttendanceHandler *handler.AttendanceHandler
	AuthMiddleware    *middleware.AuthMiddleware
}
