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
	// Membuat koneksi ke database SQL Server menggunakan fungsi dari database/sql.go
	db, err := database.NewSQLServerConnection(cfg)
	if err != nil {
		logging.Log.Fatal("failed to connect to database", zap.Error(err))
	}

	// Repository dan Service untuk User
	userRepo := user.NewSQLRepository(db)
	trxHistoryRepo := history.NewSQLRepository(db)

	trxHistoryService := history.NewTrxHistoryService(trxHistoryRepo)
	userService := user.NewService(userRepo)

	// Handler untuk User
	userHandler := handler.NewUserHandler(userService)

	// Handler lainnya
	trxHistoryHandler := handler.NewTrxHistoryHandler(trxHistoryService)

	// Inisialisasi DI Container
	return &Container{
		Config:        cfg,
		Log:           logging.Log,
		TrxHandler:    trxHistoryHandler,
		HealthHandler: handler.NewHealthHandler(),
		UserHandler:   userHandler,
		UserService:   userService,
		TrxService:    trxHistoryService,
		DB:            db, // Menyuntikkan koneksi database ke dalam container
	}
}
