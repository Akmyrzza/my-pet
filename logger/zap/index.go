package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

func NewLogger() *ZapLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil
	}

	return &ZapLogger{logger: logger, sugaredLogger: logger.Sugar()}
}

func (l *ZapLogger) Info(msg string, args ...any) {
	l.sugaredLogger.Infow(msg, args...)
}

func (l *ZapLogger) Error(msg string, args ...any) {
	l.sugaredLogger.Errorw(msg, args...)
}
