package cmd

import (
	"github.com/Akmyrzza/my-pet/internal/logger/zap"
)

func Execute() {
	logger := zap.NewZapLogger()
	logger.Info("Logger initialized")
}
