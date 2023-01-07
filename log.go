package yaglogger

import (
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"os"
	"time"
)

// Logger Struture
type Logger struct {
	Level       level
	Output      LevelOutput
	ShowLevel   bool
	ShowDate    bool
	Color       bool
	LogToScreen bool
	LogFilePath string
	LogFile     *os.File
}

// New creates a new logger
func New() *Logger {
	return &Logger{
		Level: LevelInfo,
		Output: LevelOutput{
			Fatal: os.Stderr,
			Error: os.Stderr,
			Warn:  os.Stdout,
			Info:  os.Stdout,
			Debug: os.Stdout,
			Trace: os.Stdout,
		},
		ShowDate:    true,
		ShowLevel:   true,
		Color:       true,
		LogToScreen: true,
		LogFilePath: "",
	}
}

// Log is the generic logging function
func (l *Logger) Log(level level, message string, a ...any) {
	var logDate = time.Now().Format("2006-01-02 15:04:05")

	if l.IsLogLevelEnabled(level) {
		if l.LogToScreen {
			l.screenLog(level, message, logDate, a...)
		}

		if l.LogFilePath != "" {
			if l.LogFile == nil {
				l.LogFile = l.logFileOpen()
			}
			l.fileLog(level, message, logDate, a...)
		}
	}
}

func (l *Logger) logFileOpen() *os.File {
	file, err := os.OpenFile(l.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening log file: %v", err)
		os.Exit(1)
	}
	return file
}

func (l *Logger) screenLog(level level, message, logDate string, a ...any) {
	var logLevel string
	if l.ShowDate {
		if l.Color {
			logDate = aurora.Sprintf(aurora.Bold(logDate))
		}
		if l.ShowLevel {
			logDate = logDate + " - "
		}
	} else {
		logDate = ""
	}

	if l.ShowLevel {
		logLevel = l.LevelName(level)
		if l.Color {
			logLevel = aurora.Sprintf(l.LevelColor(level, aurora.Bold(logLevel)))
		}
	}
	logline := fmt.Sprintf("[ "+logDate+logLevel+" ] "+message+"\n", a...)
	fmt.Fprint(l.LevelOutput(level), logline)
}

func (l *Logger) fileLog(level level, message, logDate string, a ...any) {
	var logLevel string
	if !l.ShowDate {
		logDate = ""
	}

	if l.ShowLevel {
		logLevel = l.LevelName(level)
	}
	fmt.Fprintf(l.LogFile, "[ "+logDate+logLevel+" ] "+message+"\n", a...)
}

// Trace logs a message with level trace
func (l *Logger) Trace(message string, a ...any) {
	l.Log(LevelTrace, message, a...)
}

// Debug logs a message with level debug
func (l *Logger) Debug(message string, a ...any) {
	l.Log(LevelDebug, message, a...)
}

// Info logs a message with level info
func (l *Logger) Info(message string, a ...any) {
	l.Log(LevelInfo, message, a...)
}

// Warn logs a message with level warn
func (l *Logger) Warn(message string, a ...any) {
	l.Log(LevelWarn, message, a...)
}

// Error logs a message with level error
func (l *Logger) Error(message string, a ...any) {
	l.Log(LevelError, message, a...)
}

// Fatal logs a message with level fatal
func (l *Logger) Fatal(message string, a ...any) {
	l.Log(LevelFatal, message, a...)
	os.Exit(1)
}
