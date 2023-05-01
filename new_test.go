package yaglogger

import (
	"github.com/logrusorgru/aurora/v4"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Logger
	}{
		{
			name: "TestNew",
			want: &Logger{
				Level: LevelInfo,
				Output: LevelOutput{
					Fatal: os.Stderr,
					Error: os.Stderr,
					Warn:  os.Stdout,
					Info:  os.Stdout,
					Debug: os.Stdout,
					Trace: os.Stdout,
				},
				Format: Format{
					ShowDate:   true,
					DateLayout: "2006-01-02 15:04:05",
					Color:      true,
					ShowLevel:  true,
					EndOfLine:  "\n",
					Prefix:     "",
				},
				LogToScreen: true,
				LogFilePath: "",
				au:          aurora.New(aurora.WithColors(true)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
