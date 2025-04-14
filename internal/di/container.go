package di

import (
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/config"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/domain/history"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/handler"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/logging"
	"go.uber.org/zap"
)

type Container struct {
	Config        *config.App
	Log           *zap.Logger
	TrxHandler    *handler.TrxHistoryHandler
	HealthHandler *handler.HealthHandler

	// @auto:inject:field
}

func InitContainer(cfg *config.App) *Container {
	// Initialize the DI container with the provided configuration

	// @auto:inject:init-service
	trxService := history.NewTrxHistoryService()
	return &Container{
		Config:        cfg,
		Log:           logging.Log,
		TrxHandler:    handler.NewTrxHistoryHandler(trxService),
		HealthHandler: handler.NewHealthHandler(),
		// @auto:inject:init-handler
	}
}
