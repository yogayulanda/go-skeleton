package di

import (
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/config"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/database"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/domain/history"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/domain/user"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/handler"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/logging"
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
	userService := user.NewService(userRepo)

	// Handler untuk User
	userHandler := handler.NewUserHandler(userService)

	// Handler lainnya
	trxService := history.NewTrxHistoryService()

	// Inisialisasi DI Container
	return &Container{
		Config:        cfg,
		Log:           logging.Log,
		TrxHandler:    handler.NewTrxHistoryHandler(trxService),
		HealthHandler: handler.NewHealthHandler(),
		UserHandler:   userHandler,
		UserService:   userService,
		DB:            db, // Menyuntikkan koneksi database ke dalam container
	}
}
