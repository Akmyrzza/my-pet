package logger

type MyLogger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}