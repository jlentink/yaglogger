package yaglogger

import (
	"github.com/logrusorgru/aurora/v4"
	"os"
)

// New creates a new logger
func New() *Logger {
	return &Logger{
		Level: LevelInfo,
		Output: LevelOutput{
			Fatal:    os.Stderr,
			Error:    os.Stderr,
			Warn:     os.Stdout,
			Info:     os.Stdout,
			Debug:    os.Stdout,
			Trace:    os.Stdout,
			Msg:      os.Stdout,
			DebugMsg: os.Stdout,
		},
		Format: Format{
			ShowDate:   true,
			DateLayout: "2006-01-02 15:04:05",
			Color:      true,
			ShowLevel:  true,
			EndOfLine:  "\n",
			Prefix:     "",
		},
		LogToScreen:  true,
		LogFilePath:  "",
		au:           aurora.New(aurora.WithColors(true)),
		ForceNewLine: true,
	}
}
