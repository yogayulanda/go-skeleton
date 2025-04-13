package di

import (
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/config"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/handler"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/logging"
	"go.uber.org/zap"
)

type Container struct {
	Config        *config.App
	Logger        *zap.Logger
	TrxHandler    *handler.TrxHistoryHandler
	HealthHandler *handler.HealthHandler
}

func InitContainer(cfg *config.App) *Container {
	logging.InitLogger(cfg.APP_MODE == "PROD")
	return &Container{
		Config:        cfg,
		Logger:        logging.Logger,
		TrxHandler:    handler.NewTrxHistoryHandler(),
		HealthHandler: handler.NewHealthHandler(),
	}
}
