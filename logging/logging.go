package logging

import (
	"context"
	"github.com/stdxgo/common-utils/logging/internal"
)

// Debug 主日志，debug 级别日志
func Debug(ctx context.Context, v ...interface{}) {
	internal.MainLogger.Debug(ctx, v...)
}

// Debugf 主日志，debug 级别日志 format打印
func Debugf(ctx context.Context, format string, v ...interface{}) {
	internal.MainLogger.Debugf(ctx, format, v...)
}

// Info 主日志，info 级别日志
func Info(ctx context.Context, v ...interface{}) {
	internal.MainLogger.Info(ctx, v...)
}

// Infof 主日志，info 级别日志 format打印
func Infof(ctx context.Context, format string, v ...interface{}) {
	internal.MainLogger.Infof(ctx, format, v...)
}

// Warn 主日志，warn 级别日志
func Warn(ctx context.Context, v ...interface{}) {
	internal.MainLogger.Warn(ctx, v...)
}

// Warnf 主日志，warn 级别日志 format打印
func Warnf(ctx context.Context, format string, v ...interface{}) {
	internal.MainLogger.Warnf(ctx, format, v...)
}

// Error 主日志，error 级别日志
func Error(ctx context.Context, v ...interface{}) {
	internal.MainLogger.Error(ctx, v...)
}

// Errorf 主日志，error 级别日志 format打印
func Errorf(ctx context.Context, format string, v ...interface{}) {
	internal.MainLogger.Errorf(ctx, format, v...)
}

// Fatal 主日志，fatal 级别日志
func Fatal(ctx context.Context, v ...interface{}) {
	internal.MainLogger.Fatal(ctx, v...)
}

// Fatalf 主日志，fatal 级别日志 format打印
func Fatalf(ctx context.Context, format string, v ...interface{}) {
	internal.MainLogger.Fatalf(ctx, format, v...)
}
