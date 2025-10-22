package logger

type Logger interface {
	Info(msg string)
	Error(msg string)
	Fatal(msg string)

	// Structured logging
	WithField(key string, value any) Logger
	WithFields(fields map[string]any) Logger
}
