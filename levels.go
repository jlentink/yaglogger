package yaglogger

import (
	"io"
)

// Level type
type LogLevel uint

// All log levels
const (
	LevelFatal LogLevel = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
	LevelNone
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
	Fatal    io.Writer
	Error    io.Writer
	Warn     io.Writer
	Info     io.Writer
	Debug    io.Writer
	Trace    io.Writer
	Msg      io.Writer
	DebugMsg io.Writer
}
