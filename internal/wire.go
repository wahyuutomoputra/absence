//go:build wireinject
// +build wireinject

package internal

import (
	"absence/internal/handler"
	"absence/internal/repository"
	"absence/internal/service"
	"absence/pkg/database"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewAttendanceRepository,
	service.NewUserService,
	service.NewAttendanceService,
	handler.NewUserHandler,
	handler.NewAttendanceHandler,
)

type API struct {
	UserHandler       *handler.UserHandler
	AttendanceHandler *handler.AttendanceHandler
}

func InitializeAPI(db *gorm.DB) (*API, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(API), "*"),
	)
	return nil, nil
}

func InitializeDB(config *database.Config) (*gorm.DB, error) {
	return database.NewDatabase(config)
}
