package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger(appMode string) {
	var err error
	if appMode == "PROD" {
		Log, err = zap.NewProduction()
	} else {
		cfg := zap.NewDevelopmentConfig()

		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stderr"}

		Log, err = cfg.Build()
	}

	if err != nil {
		panic("unable to initialize zap logger: " + err.Error())
	}

	// ⛳ Replace global zap logger so zap.L() also uses this config
	zap.ReplaceGlobals(Log)
}

// SyncLogger flushes any buffered log entries
func SyncLogger() {
	if Log != nil {
		_ = Log.Sync()
	}
}
