package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger(isProduction bool) {
	var err error
	if isProduction {
		Logger, err = zap.NewProduction()
	} else {
		Logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("unable to initialize zap logger: " + err.Error())
	}
}

// SyncLogger flushes any buffered log entries
func SyncLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
