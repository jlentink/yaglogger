package yaglogger

import (
	"testing"
)

func TestNew(t *testing.T) {
	logger := New()

	// Verify Level field
	if logger.Level != LevelInfo {
		t.Errorf("Expected Level to be LevelInfo, got %v", logger.Level)
	}

	//// Verify Output fields
	//expectedOutputs := map[LogLevel]*os.File{
	//	Fatal:    os.Stderr,
	//	Error:    os.Stderr,
	//	Warn:     os.Stdout,
	//	Info:     os.Stdout,
	//	Debug:    os.Stdout,
	//	Trace:    os.Stdout,
	//	Msg:      os.Stdout,
	//	DebugMsg: os.Stdout,
	//}
	//
	//for level, expectedOutput := range expectedOutputs {
	//	if logger.Output[level] != expectedOutput {
	//		t.Errorf("Expected Output[%v] to be %v, got %v", level, expectedOutput, logger.Output[level])
	//	}
	//}

	// Verify Format fields
	if !logger.Format.ShowDate {
		t.Error("Expected ShowDate to be true")
	}

	if logger.Format.DateLayout != "2006-01-02 15:04:05" {
		t.Errorf("Expected DateLayout to be \"2006-01-02 15:04:05\", got %v", logger.Format.DateLayout)
	}

	if !logger.Format.Color {
		t.Error("Expected Color to be true")
	}

	if !logger.Format.ShowLevel {
		t.Error("Expected ShowLevel to be true")
	}

	if logger.Format.EndOfLine != "\n" {
		t.Errorf("Expected EndOfLine to be \"\\n\", got %v", logger.Format.EndOfLine)
	}

	if logger.Format.Prefix != "" {
		t.Error("Expected Prefix to be an empty string")
	}

	// Verify other fields
	if !logger.LogToScreen {
		t.Error("Expected LogToScreen to be true")
	}

	if logger.LogFilePath != "" {
		t.Error("Expected LogFilePath to be an empty string")
	}

	if logger.au == nil {
		t.Error("Expected au to be initialized")
	}

	if !logger.ForceNewLine {
		t.Error("Expected ForceNewLine to be true")
	}
}
