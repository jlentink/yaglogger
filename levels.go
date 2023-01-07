package yaglogger

import (
	"github.com/logrusorgru/aurora/v3"
	"io"
	"os"
)

// Level type
type level uint

// All log levels
const (
	LevelFatal level = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
)

// LevelOutput defines the output stream for screen
type LevelOutput struct {
	Fatal io.Writer
	Error io.Writer
	Warn  io.Writer
	Info  io.Writer
	Debug io.Writer
	Trace io.Writer
}

// LevelOutput returns the output stream for the given level
func (l *Logger) LevelOutput(level level) io.Writer {
	switch level {
	case LevelFatal:
		return l.Output.Fatal
	case LevelError:
		return l.Output.Error
	case LevelWarn:
		return l.Output.Warn
	case LevelInfo:
		return l.Output.Info
	case LevelDebug:
		return l.Output.Debug
	case LevelTrace:
		return l.Output.Trace
	default:
		return os.Stdout
	}
}

// LevelName returns the name of the given level
func (l *Logger) LevelName(level level) string {
	switch level {
	case LevelFatal:
		return "  FATAL  "
	case LevelError:
		return "  ERROR  "
	case LevelWarn:
		return " WARNING "
	case LevelInfo:
		return "  INFO   "
	case LevelDebug:
		return "  DEBUG  "
	case LevelTrace:
		return "  TRACE  "
	default:
		return " UNKNOWN "
	}
}

// LevelColor returns the color of the given level
func (l *Logger) LevelColor(level level, message any) aurora.Value {
	switch level {
	case LevelFatal:
		return aurora.Red(message)
	case LevelError:
		return aurora.Magenta(message)
	case LevelWarn:
		return aurora.Yellow(message)
	case LevelInfo:
		return aurora.Green(message)
	case LevelDebug:
		return aurora.Cyan(message)
	case LevelTrace:
		return aurora.White(message)
	default:
		return aurora.White(message)
	}
}

// IsLogLevelEnabled returns true if the given level is enabled
func (l *Logger) IsLogLevelEnabled(level level) bool {
	return l.Level >= level
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level level) {
	if level >= LevelFatal && level <= LevelAll {
		l.Level = level
	}
}
