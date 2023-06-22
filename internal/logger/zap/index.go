package zap

import (
	"github.com/Akmyrzza/my-pet/internal/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

func NewZapLogger() logger.MyLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, _ := config.Build(zap.AddCallerSkip(1))

	return &ZapLogger{logger: logger, sugaredLogger: logger.Sugar()}
}

func (ZapLogger *ZapLogger) Info(msg string, args ...any) {
	ZapLogger.sugaredLogger.Infow(msg, args...)
}

func (ZapLogger *ZapLogger) Error(msg string, args ...any) {
	ZapLogger.sugaredLogger.Errorw(msg, args...)
}
