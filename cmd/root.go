package cmd

import (
	"github.com/Akmyrzza/my-pet/logger/zap"
)

func Execute() {

	logger := zap.NewLogger()
	logger.Info("Logger initialized")
}
