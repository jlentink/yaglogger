package yaglogger

import (
	"io"
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
	Fatal    io.Writer
	Error    io.Writer
	Warn     io.Writer
	Info     io.Writer
	Debug    io.Writer
	Trace    io.Writer
	Msg      io.Writer
	DebugMsg io.Writer
}
