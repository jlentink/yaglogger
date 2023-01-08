package yaglogger

var (
	_defaultLogger = New()
)

// SetLevel sets the log level
//
//goland:noinspection GoUnusedExportedFunction
func SetLevel(level level) {
	_defaultLogger.SetLevel(level)
}

// Trace logs a message with level trace
//
//goland:noinspection GoUnusedExportedFunction
func Trace(message string, a ...any) {
	_defaultLogger.Trace(message, a...)
}

// Debug logs a message with level debug
//
//goland:noinspection GoUnusedExportedFunction
func Debug(message string, a ...any) {
	_defaultLogger.Debug(message, a...)
}

// Info logs a message with level info
//
//goland:noinspection GoUnusedExportedFunction
func Info(message string, a ...any) {
	_defaultLogger.Info(message, a...)
}

// Warn logs a message with level warn
//
//goland:noinspection GoUnusedExportedFunction
func Warn(message string, a ...any) {
	_defaultLogger.Warn(message, a...)
}

// Error logs a message with level error
//
//goland:noinspection GoUnusedExportedFunction
func Error(message string, a ...any) {
	_defaultLogger.Error(message, a...)
}

// Fatal logs a message with level fatal
//
//goland:noinspection GoUnusedExportedFunction
func Fatal(message string, a ...any) {
	_defaultLogger.Fatal(message, a...)
}

// Fatalf logs a message with level fatal
//
//goland:noinspection GoUnusedExportedFunction
func Fatalf(message string, a ...any) {
	_defaultLogger.Fatal(message, a...)
}

// Panic logs a message with level panic
//
//goland:noinspection GoUnusedExportedFunction
func Panic(message string, a ...any) {
	_defaultLogger.Panic(message, a...)
}

// Panicf logs a message with level panic
//
//goland:noinspection GoUnusedExportedFunction
func Panicf(message string, a ...any) {
	_defaultLogger.Panic(message, a...)
}
