package logr

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config holds the configuration for the logger.
// It includes the log level, which can be set to "debug", "info", "warn", "error", or "fatal".
type Config struct {
	Level string `json:"level,omitempty" koanf:"level"`
}

// New creates a zap logger for console.
func New(cfg Config) (*zap.Logger, error) {
	var lvl zapcore.Level

	err := lvl.Set(cfg.Level)
	if err != nil {
		return nil, err
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	defaultCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stderr)), lvl)
	cores := []zapcore.Core{
		defaultCore,
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return logger, nil
}
