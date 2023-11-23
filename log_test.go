package yaglogger

import (
	"github.com/logrusorgru/aurora/v4"
	"io"
	"testing"
)

func TestLogger_SetLevelByString(t *testing.T) {
	type fields struct {
		Level        LogLevel
		Output       LevelOutput
		Format       Format
		LogToScreen  bool
		LogFilePath  string
		LogFile      io.Writer
		au           *aurora.Aurora
		ForceNewLine bool
	}
	type args struct {
		level string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   LogLevel
	}{
		{
			name: "TestLogger_SetLevelByString - fatal",
			args: args{level: "fatal"},
			want: LevelFatal,
		},
		{
			name: "TestLogger_SetLevelByString - error",
			args: args{level: "error"},
			want: LevelError,
		},
		{
			name: "TestLogger_SetLevelByString - warn",
			args: args{level: "Warn"},
			want: LevelWarn,
		},
		{
			name: "TestLogger_SetLevelByString - warning",
			args: args{level: "Warning"},
			want: LevelWarn,
		},
		{
			name: "TestLogger_SetLevelByString - info",
			args: args{level: "debug"},
			want: LevelDebug,
		},
		{
			name: "TestLogger_SetLevelByString - trace",
			args: args{level: "trace"},
			want: LevelTrace,
		},
		{
			name: "TestLogger_SetLevelByString - all",
			args: args{level: "all"},
			want: LevelAll,
		},
		{
			name: "TestLogger_SetLevelByString - none",
			args: args{level: "none"},
			want: LevelNone,
		},
		{
			name: "TestLogger_SetLevelByString - default",
			args: args{level: "ssaa"},
			want: LevelAll,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			if got := l.SetLevelByString(tt.args.level); got != tt.want {
				t.Errorf("SetLevelByString() = %v, want %v", got, tt.want)
			}
		})
	}
}
