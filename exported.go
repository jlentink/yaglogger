package yaglogger

import (
	"io"
	"log"
)

var (
	_defaultLogger = New()
)

func GetInstance() *Logger {
	return _defaultLogger
}

func SetLogger(logger *Logger) {
	_defaultLogger = logger
}

// SetLevel sets the log level
func SetLevel(level LogLevel) {
	log.SetFlags(0)
	_defaultLogger.SetLevel(level)
}

func SetLevelByString(level string) {
	_defaultLogger.SetLevelByString(level)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...any) {
	_defaultLogger.Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...any) {
	_defaultLogger.Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...any) {
	_defaultLogger.Fatalln(v...)
}

// Flags returns the output flags for the standard logger. and is unsupported in this package.
func Flags() int {
	return int(_defaultLogger.Level)
}

func Output(calldepth int, s string) error {
	panic("not implemented")
}

func Panic(v ...any) {
	_defaultLogger.Panic(v...)
}

func Panicf(format string, v ...any) {
	_defaultLogger.Panicf(format, v...)
}

func Panicln(v ...any) {
	_defaultLogger.Panic(v...)
}

func Prefix() string {
	panic("not implemented")
}

func Print(v ...any) {
	_defaultLogger.Print(v...)
}

func Printf(format string, v ...any) {
	_defaultLogger.Printf(format, v...)
}

func Println(v ...any) {
	panic("not implemented")
}

func SetFlags(flag int) {
	panic("not implemented")
}

func SetOutput(w io.Writer) {
	_defaultLogger.LogFile = w
}
func SetPrefix(prefix string) {
	panic("not implemented")
}

func Writer() io.Writer {
	return _defaultLogger.LogFile
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
