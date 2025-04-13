package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(isProduction bool) {
	var err error
	if isProduction {
		Logger, err = zap.NewProduction()
	} else {
		cfg := zap.NewDevelopmentConfig()

		// Colorize log levels (INFO, ERROR, etc.)
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

		// Set output to stdout
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stderr"}

		Logger, err = cfg.Build()
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

// Replace global zap logger with our instance (optional, useful for libraries that use zap.L())
func ReplaceGlobals() {
	if Logger != nil {
		zap.ReplaceGlobals(Logger)
	}
}
