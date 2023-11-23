package yaglogger

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

// Logger Struture
type Logger struct {
	Level        LogLevel
	Output       LevelOutput
	Format       Format
	LogToScreen  bool
	LogFilePath  string
	LogFile      io.Writer
	au           *aurora.Aurora
	ForceNewLine bool
}

func (l *Logger) SetLevelByString(level string) LogLevel {
	switch strings.ToLower(level) {
	case "fatal":
		l.Level = LevelFatal
	case "error":
		l.Level = LevelError
	case "warn", "warning":
		l.Level = LevelWarn
	case "info":
		l.Level = LevelInfo
	case "debug":
		l.Level = LevelDebug
	case "trace":
		l.Level = LevelTrace
	case "all":
		l.Level = LevelAll
	case "none":
		l.Level = LevelNone
	default:
		l.Level = LevelAll
		l.Error("Unknown log level: %s", level)
	}
	return l.Level
}

// EnableColors enables/disables colors
func (l *Logger) EnableColors(c bool) {
	l.Format.Color = c
	l.au = aurora.New(aurora.WithColors(c))
}

// GetColours returns the aurora instance
func (l *Logger) GetColours() *aurora.Aurora {
	return l.au
}

// Log is the generic logging function
func (l *Logger) Log(level LogLevel, message string, a ...any) {
	var logDate = time.Now().Format(l.Format.DateLayout)

	message = strings.TrimSuffix(message, "\n")
	message = strings.TrimSuffix(message, "\r")

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

func (l *Logger) logFileOpen() io.Writer {
	file, err := os.OpenFile(l.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//nolint:errcheck
		//goland:noinspection GoUnhandledErrorResult
		fmt.Fprintf(os.Stderr, "Error opening log file: %v", err)
		os.Exit(1)
	}
	return file
}

func (l *Logger) screenLog(level LogLevel, message, logDate string, a ...any) {
	var logLevel string
	if l.Format.ShowDate {
		if l.Format.Color {
			logDate = aurora.Sprintf(aurora.Bold(logDate))
		}
		if l.Format.ShowLevel {
			logDate = logDate + " - "
		}
	} else {
		logDate = ""
	}

	if l.Format.ShowLevel {
		logLevel = l.LevelName(level)
		if l.Format.Color {
			logLevel = aurora.Sprintf(l.LevelColor(level, aurora.Bold(logLevel)))
		}
	}
	logLine := fmt.Sprintf("[ "+logDate+logLevel+" ] "+message+l.Format.EndOfLine, a...)
	_, _ = l.Fprint(l.LevelOutput(level), logLine)
}

func (l *Logger) fileLog(level LogLevel, message, logDate string, a ...any) {
	var logLevel string
	if !l.Format.ShowDate {
		logDate = ""
	}

	if l.Format.ShowLevel {
		logLevel = l.LevelName(level)
	}
	//goland:noinspection GoUnhandledErrorResult
	fmt.Fprintf(l.LogFile, "[ "+logDate+logLevel+" ] "+message+l.Format.EndOfLine, a...)
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
func (l *Logger) Fatal(v ...any) {
	message := fmt.Sprint(v...)
	l.Log(LevelFatal, message)
	os.Exit(1)
}

func (l *Logger) Fatalln(v ...any) {
	message := fmt.Sprintln(v...)
	l.Log(LevelFatal, message)
	os.Exit(1)
}

// Fatalf logs a message with level fatal
func (l *Logger) Fatalf(message string, a ...any) {
	l.Log(LevelFatal, message, a...)
	os.Exit(1)
}

// Panic logs a message with level fatal
func (l *Logger) Panic(v ...any) {
	message := fmt.Sprint(v...)
	l.Log(LevelFatal, message)
	panic(l.newLine(fmt.Sprint(message)))
}

// Panic logs a message with level fatal
func (l *Logger) Panicf(message string, a ...any) {
	l.Log(LevelFatal, message, a...)
	panic(l.newLine(fmt.Sprintf(message, a...)))
}

// LevelOutput returns the output stream for the given level
func (l *Logger) LevelOutput(level LogLevel) io.Writer {
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
func (l *Logger) LevelName(level LogLevel) string {
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
func (l *Logger) LevelColor(level LogLevel, message any) aurora.Value {
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
func (l *Logger) IsLogLevelEnabled(level LogLevel) bool {
	return l.Level >= level
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) *Logger {
	if level >= LevelFatal && level <= LevelAll {
		l.Level = level
	}
	return l
}

// SetDebug sets the log level to debug
func (l *Logger) SetDebug(d bool) *Logger {
	if d {
		l.Level = LevelDebug
	} else {
		l.Level = LevelInfo
	}
	return l
}

func (l *Logger) isInstanceOf(object, objectType interface{}) bool {
	return reflect.TypeOf(object) == reflect.TypeOf(objectType)
}

func (l *Logger) newLine(in string) string {
	if l.ForceNewLine && !strings.HasSuffix(in, l.Format.EndOfLine) {
		return in + l.Format.EndOfLine
	}
	return in
}

// Print prints message
func (l *Logger) Print(v ...any) {
	l.Fprint(l.Output.Msg, fmt.Sprint(v...))
}

// Printf prints message with formatting
func (l *Logger) Printf(format string, a ...any) {
	l.Fprintf(l.Output.Msg, format, a...)
}

// PrintDebug prints message with debug level
func (l *Logger) PrintDebug(format any) {
	if l.IsLogLevelEnabled(LevelDebug) {
		if l.isInstanceOf(format, aurora.Value{}) {
			format = aurora.Sprintf("%s", format)
		}
		l.Fprint(l.Output.Msg, format)
	}
}

// PrintDebugf prints message with debug level and formatting
func (l *Logger) PrintDebugf(format string, a ...any) {
	if l.IsLogLevelEnabled(LevelDebug) {
		l.Fprintf(l.Output.Msg, format, a...)
	}
}

func (l *Logger) Fprint(w io.Writer, a ...any) (n int, err error) {
	format := l.newLine(fmt.Sprint(a...))
	return fmt.Fprint(w, format)
}

func (l *Logger) Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	format = l.newLine(format)
	return fmt.Fprintf(w, format, a...)
}
