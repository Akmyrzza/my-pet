package zap

import "go.uber.org/zap"

type ZapLogger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

func NewLogger() error {

}

func (l *ZapLogger) Info(msg string, args ...any) {
	l.sugaredLogger.Infow(msg, args...)
}

func (l *ZapLogger) Error(msg string, args ...any) {
	l.sugaredLogger.Errorw(msg, args...)
}
