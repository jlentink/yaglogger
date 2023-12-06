package yaglogger

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"github.com/mattn/go-isatty"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// Logger Structure
type Logger struct {
	Level           LogLevel
	Output          LevelOutput
	Format          Format
	LogToScreen     bool
	LogFilePath     string
	LogFile         io.Writer
	au              *aurora.Aurora
	ForceNewLine    bool
	ShowLogLocation bool
}

func (l *Logger) SetLevelByString(level string) LogLevel {
	switch strings.ToLower(level) {
	case LevelFatal.String():
		l.Level = LevelFatal
	case LevelError.String():
		l.Level = LevelError
	case LevelWarn.String(), "warning":
		l.Level = LevelWarn
	case LevelInfo.String():
		l.Level = LevelInfo
	case LevelDebug.String():
		l.Level = LevelDebug
	case LevelTrace.String():
		l.Level = LevelTrace
	case LevelAll.String():
		l.Level = LevelAll
	case LevelNone.String():
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

func (l *Logger) GetFileWithLine() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && !strings.Contains(file, "yaglogger") || ok && strings.Contains(file, "main.go") {
			return fmt.Sprintf("%s:%d", file, line)
		}
		//fmt.Println(file, line)
	}
	return ""
}

// Log is the generic logging function
func (l *Logger) Log(level LogLevel, message string, a ...any) {
	var logDate = time.Now().Format(l.Format.DateLayout)
	location := ""
	if l.ShowLogLocation {
		location = l.GetFileWithLine()
	}

	message = strings.TrimSuffix(message, "\n")
	message = strings.TrimSuffix(message, "\r")

	if l.IsLogLevelEnabled(level) {
		if l.LogToScreen {
			l.screenLog(level, message, logDate, location, a...)
		}

		if l.LogFilePath != "" {
			if l.LogFile == nil {
				l.LogFile = l.logFileOpen()
			}
			l.fileLog(level, message, logDate, location, a...)
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

func (l *Logger) screenLog(level LogLevel, message, logDate, location string, a ...any) {
	var logLevel string
	var useColor = l.Format.Color

	if !isatty.IsTerminal(l.LevelOutput(level).Fd()) {
		useColor = false

	}

	if l.Format.ShowDate {
		if useColor {
			logDate = aurora.Sprintf(aurora.Bold(logDate))
		}
		if l.Format.ShowLevel {
			logDate = logDate + " - "
		}
	} else {
		logDate = ""
	}

	if l.Format.ShowLevel {
		logLevel = l.CenteredLevelName(level)
		if useColor {
			logLevel = aurora.Sprintf(l.LevelColor(level, aurora.Bold(logLevel)))
		}
	}
	if location != "" {
		location = " - " + location + " "
	}
	logLine := fmt.Sprintf("[ "+logDate+logLevel+" "+location+"] "+message+l.Format.EndOfLine, a...)
	_, _ = l.Fprint(l.LevelOutput(level), logLine)
}

func (l *Logger) fileLog(level LogLevel, message, logDate, location string, a ...any) {
	var logLevel string
	if !l.Format.ShowDate {
		logDate = ""
	}

	if l.Format.ShowLevel {
		logLevel = l.CenteredLevelName(level)
	}
	if location != "" {
		location = " - " + location + " "
	}
	//goland:noinspection GoUnhandledErrorResult
	fmt.Fprintf(l.LogFile, "[ "+logDate+logLevel+" "+location+"] "+message+l.Format.EndOfLine, a...)
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
func (l *Logger) LevelOutput(level LogLevel) *os.File {
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

// CenteredLevelName returns the name of the given level
func (l *Logger) CenteredLevelName(level LogLevel) string {
	levelName := strings.ToUpper(level.String())
	width := 10
	padding := strings.Repeat(" ", (width-len(levelName))/2)
	formatted := fmt.Sprintf("%s%s%s", padding, levelName, padding)
	return formatted[0 : width-1]
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
	//goland:noinspection GoUnhandledErrorResult
	l.Fprint(l.Output.Msg, fmt.Sprint(v...))
}

// Printf prints message with formatting
func (l *Logger) Printf(format string, a ...any) {
	//goland:noinspection GoUnhandledErrorResult
	l.Fprintf(l.Output.Msg, format, a...)
}

// PrintDebug prints message with debug level
func (l *Logger) PrintDebug(format any) {
	if l.IsLogLevelEnabled(LevelDebug) {
		if l.isInstanceOf(format, aurora.Value{}) {
			format = aurora.Sprintf("%s", format)
		}
		//goland:noinspection GoUnhandledErrorResult
		l.Fprint(l.Output.Msg, format)
	}
}

// PrintDebugf prints message with debug level and formatting
func (l *Logger) PrintDebugf(format string, a ...any) {
	if l.IsLogLevelEnabled(LevelDebug) {
		//goland:noinspection GoUnhandledErrorResult
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
