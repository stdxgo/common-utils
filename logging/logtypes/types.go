package logtypes

import (
	"context"
	"time"
)

// LogConfig logConfig
type LogConfig struct {
	AppName  string        `json:"logFile"`
	LogPath  string        `json:"logPath"`
	LogAge   time.Duration `json:"logAge"`
	LogLevel string        `json:"logLevel"` // debug info warn error fatal
}

// Level 日志级别
type Level int

// LevelFlag 日志级别标识
type LevelFlag string

// LogFormatFunc 输出日志格式化func
type LogFormatFunc func(ctx context.Context, flag LevelFlag, callDeep int, tid, body string) string

const (
	// DebugLv 日志级别
	DebugLv Level = 0 + iota
	// InfoLv  日志级别
	InfoLv
	// WarnLv 日志级别
	WarnLv
	// ErrorLv 日志级别
	ErrorLv
	// FatalLv 日志级别
	FatalLv

	// DEBUG 日志级别标识
	DEBUG LevelFlag = "DEBUG"
	// INFO 日志级别标识
	INFO LevelFlag = "INFO"
	// WARN 日志级别标识
	WARN LevelFlag = "WARN"
	// ERROR 日志级别标识
	ERROR LevelFlag = "ERROR"
	// FATAL 日志级别标识
	FATAL LevelFlag = "FATAL"
)

// AppLogger appLogger
type AppLogger interface {
	Debug(ctx context.Context, v ...interface{})
	Debugf(ctx context.Context, format string, v ...interface{})
	Info(ctx context.Context, v ...interface{})
	Infof(ctx context.Context, format string, v ...interface{})
	Warn(ctx context.Context, v ...interface{})
	Warnf(ctx context.Context, format string, v ...interface{})
	Error(ctx context.Context, v ...interface{})
	Errorf(ctx context.Context, format string, v ...interface{})
	Fatal(ctx context.Context, v ...interface{})
	Fatalf(ctx context.Context, format string, v ...interface{})
}
