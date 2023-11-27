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

func TestLogger_CenteredLevelName(t *testing.T) {
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
		level LogLevel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "TestLogger_CenteredLevelName - fatal",
			args: args{level: LevelFatal},
			want: "  FATAL  ",
		},
		{
			name: "TestLogger_CenteredLevelName - error",
			args: args{level: LevelError},
			want: "  ERROR  ",
		},
		{
			name: "TestLogger_CenteredLevelName - info",
			args: args{level: LevelInfo},
			want: "   INFO  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Level:        tt.fields.Level,
				Output:       tt.fields.Output,
				Format:       tt.fields.Format,
				LogToScreen:  tt.fields.LogToScreen,
				LogFilePath:  tt.fields.LogFilePath,
				LogFile:      tt.fields.LogFile,
				au:           tt.fields.au,
				ForceNewLine: tt.fields.ForceNewLine,
			}
			if got := l.CenteredLevelName(tt.args.level); got != tt.want {
				t.Errorf("CenteredLevelName() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestLogger_EnableColors(t *testing.T) {
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
		c  bool
		au bool
	}
	type want struct {
		color bool
		au    *aurora.Aurora
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "TestLogger_EnableColors - true",
			args: args{
				c: true,
			},
			want: want{
				color: true,
				au:    aurora.New(aurora.WithColors(true)),
			},
		},
		{
			name: "TestLogger_EnableColors - false",
			args: args{
				c: false,
			},
			want: want{
				color: false,
				au:    aurora.New(aurora.WithColors(false)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			l.EnableColors(tt.args.c)
			if l.au.Config().Colors != tt.want.au.Config().Colors {
				t.Errorf("EnableColors() [au] = %v, want %v", l.au.Config().Colors, tt.want.au.Config().Colors)
			}

			if tt.want.color != l.Format.Color {
				t.Errorf("EnableColors() [color] = %v, want %v", l.Format.Color, tt.want.color)
			}
		})
	}
}
