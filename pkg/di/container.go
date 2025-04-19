package di

import (
	"github.com/yogayulanda/go-skeleton/pkg/config"
	logging "github.com/yogayulanda/go-skeleton/pkg/logger"
	"github.com/yogayulanda/go-skeleton/pkg/repository"
	"github.com/yogayulanda/go-skeleton/pkg/service"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Menggunakan var untuk deklarasi container DI global
var container *Container

type Container struct {
	Config             *config.App
	Log                *zap.Logger
	DB                 *gorm.DB // Koneksi database SQL Server
	HealthCheckService *service.HealthCheckService
	UserService        *service.UserService
}

func InitContainer(cfg *config.App) *Container {
	// Inisialisasi logger
	logging.InitLogger(cfg.LOG_LEVEL)
	log := logging.Log
	// Mengambil logger dari package logging
	// Membuat koneksi ke database SQL Server menggunakan fungsi dari database/sql.go
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}

	// init Repository dan Service
	healthCheckService := service.NewHealthCheckService(db, log)

	userRepo := repository.NewUserRepository(db, log)
	userService := service.NewUserService(userRepo, log)

	container = &Container{
		Config:             cfg,
		Log:                log,
		HealthCheckService: healthCheckService,
		UserService:        userService,
		DB:                 db, // Menyuntikkan koneksi database ke dalam container
	}
	// Inisialisasi DI Container
	return container

}
