package di

import (
	"github.com/yogayulanda/go-skeleton/internal/config"
	"github.com/yogayulanda/go-skeleton/internal/database"
	"github.com/yogayulanda/go-skeleton/internal/domain/history"
	"github.com/yogayulanda/go-skeleton/internal/domain/user"
	"github.com/yogayulanda/go-skeleton/internal/handler"
	"github.com/yogayulanda/go-skeleton/internal/logging"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Menggunakan var untuk deklarasi container DI global
var container *Container

type Container struct {
	Config        *config.App
	Log           *zap.Logger
	TrxHandler    *handler.TrxHistoryHandler
	HealthHandler *handler.HealthHandler
	UserHandler   *handler.UserHandler
	UserService   *user.Service
	TrxService    *history.TrxHistoryService
	DB            *gorm.DB // Koneksi database SQL Server
}

func InitContainer(cfg *config.App) *Container {
	// Inisialisasi logger
	logging.InitLogger(cfg.LOG_LEVEL)
	defer logging.SyncLogger()
	// Mengambil logger dari package logging
	// Membuat koneksi ke database SQL Server menggunakan fungsi dari database/sql.go
	db, err := database.NewSQLServerConnection(cfg)
	if err != nil {
		logging.Log.Fatal("failed to connect to database", zap.Error(err))
	}

	// init Repository dan Service
	userRepo := user.NewSQLRepository(db)
	userService := user.NewService(userRepo)

	trxHistoryRepo := history.NewSQLRepository(db)
	trxHistoryService := history.NewTrxHistoryService(trxHistoryRepo)

	// Handler
	userHandler := handler.NewUserHandler(userService)
	trxHistoryHandler := handler.NewTrxHistoryHandler(trxHistoryService)
	healthHandler := handler.NewHealthHandler()

	container = &Container{
		Config:        cfg,
		Log:           logging.Log,
		TrxHandler:    trxHistoryHandler,
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
		UserService:   userService,
		TrxService:    trxHistoryService,
		DB:            db, // Menyuntikkan koneksi database ke dalam container
	}

	// Inisialisasi DI Container
	return container

}
