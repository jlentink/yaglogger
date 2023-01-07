package yaglogger

var (
	_defaultLogger = New()
)

// SetLevel sets the log level
func SetLevel(level level) {
	_defaultLogger.SetLevel(level)
}

// Trace logs a message with level trace
func Trace(message string, a ...any) {
	_defaultLogger.Trace(message, a...)
}

// Debug logs a message with level debug
func Debug(message string, a ...any) {
	_defaultLogger.Debug(message, a...)
}

// Info logs a message with level info
func Info(message string, a ...any) {
	_defaultLogger.Info(message, a...)
}

// Warn logs a message with level warn
func Warn(message string, a ...any) {
	_defaultLogger.Warn(message, a...)
}

// Error logs a message with level error
func Error(message string, a ...any) {
	_defaultLogger.Error(message, a...)
}

// Fatal logs a message with level fatal
func Fatal(message string, a ...any) {
	_defaultLogger.Fatal(message, a...)
}
