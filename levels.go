package yaglogger

import (
	"os"
)

// Level type
type LogLevel uint

// All log levels
const (
	LevelNone LogLevel = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
)

var levelNames = map[LogLevel]string{
	LevelFatal: "fatal",
	LevelError: "error",
	LevelWarn:  "warn",
	LevelInfo:  "info",
	LevelDebug: "debug",
	LevelTrace: "trace",
	LevelAll:   "all",
	LevelNone:  "none",
}

func (l LogLevel) String() string {
	return levelNames[l]
}

// LevelOutput defines the output stream for screen
type LevelOutput struct {
	Fatal    *os.File
	Error    *os.File
	Warn     *os.File
	Info     *os.File
	Debug    *os.File
	Trace    *os.File
	Msg      *os.File
	DebugMsg *os.File
}
