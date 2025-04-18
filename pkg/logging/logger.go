package logging

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// initDevelopmentLogger menginisialisasi logger untuk development environment.
func initDevelopmentLogger(zapLevel zapcore.Level) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level.SetLevel(zapLevel)
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Gunakan zapcore.StackEncoder untuk stack trace
	// cfg.EncoderConfig.EncodeStack = zapcore.StackEncoder // Ini akan mencatat stack trace dalam format yang mudah dibaca

	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}

	return cfg.Build()
}

// initProductionLogger menginisialisasi logger untuk production environment.
func initProductionLogger(zapLevel zapcore.Level) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zapLevel)
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Gunakan zapcore.StackEncoder untuk stack trace
	// cfg.EncoderConfig.EncodeStack = zapcore.StackEncoder // Ini akan mencatat stack trace dalam format yang mudah dibaca

	return cfg.Build()
}

// InitLogger menginisialisasi logger berdasarkan logLevel yang diberikan dalam parameter.
func InitLogger(logLevel string) {
	var err error
	var zapLevel zapcore.Level

	// Menentukan log level yang sesuai berdasarkan parameter logLevel
	switch strings.ToLower(logLevel) {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.InfoLevel // Default ke Info jika tidak dikenali
	}

	// Pilih logger yang sesuai berdasarkan log level
	if zapLevel == zapcore.DebugLevel {
		// Development logger dengan konfigurasi berwarna dan detail
		Log, err = initDevelopmentLogger(zapLevel)
	} else {
		// Production logger dengan konfigurasi yang lebih efisien
		Log, err = initProductionLogger(zapLevel)
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
