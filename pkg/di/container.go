package di

import (
	"context"

	redisClient "github.com/redis/go-redis/v9"
	"github.com/yogayulanda/go-skeleton/pkg/config"
	"github.com/yogayulanda/go-skeleton/pkg/event"
	logging "github.com/yogayulanda/go-skeleton/pkg/logger"
	"github.com/yogayulanda/go-skeleton/pkg/redis"
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
	Redist             *redisClient.Client
	HealthCheckService *service.HealthCheckService
	UserService        *service.UserService
	ErrorRepo          *repository.ErrorCodeRepository
}

func InitContainer(cfg *config.App) *Container {
	// Inisialisasi logger
	logging.InitLogger(cfg.LOG_LEVEL)
	log := logging.Log
	// Mengambil logger dari package logging
	// Membuat koneksi ke database SQL Server menggunakan fungsi dari database/sql.go
	db, err := config.InitDB(cfg, log)
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}

	// Init Redis
	redisClient, err := config.InitRedist(cfg, log)
	if err != nil {
		log.Fatal("failed to connect to Redis", zap.Error(err))
	}

	// Init Error Cache (Singleton) menggunakan InitErrorCache
	errorCache := redis.InitErrorCache(redisClient, log)

	// Init Error Code Repository
	errorCodeRepo := repository.NewErrorCodeRepository(db)

	// Start listening for error code updates
	errorCodeEvent := event.NewErrorCodeEvent(errorCache, errorCodeRepo, log)
	go errorCodeEvent.SubscribeErrorCodeChanges(context.Background(), redisClient, "error_code_updates")

	// init Repository dan Service
	healthCheckService := service.NewHealthCheckService(db, log)

	userRepo := repository.NewUserRepository(db, log)
	errorRepo := repository.NewErrorCodeRepository(db)
	userService := service.NewUserService(userRepo, log)

	container = &Container{
		Config:             cfg,
		Log:                log,
		HealthCheckService: healthCheckService,
		UserService:        userService,
		DB:                 db, // Menyuntikkan koneksi database ke dalam container
		Redist:             redisClient,
		ErrorRepo:          errorRepo,
	}
	// Inisialisasi DI Container
	return container

}
