package internal

import (
	"context"
	"fmt"
	"github.com/stdxgo/common-utils/excontext"
	"github.com/stdxgo/common-utils/exruntime"
	"github.com/stdxgo/common-utils/logging/logtypes"
	"os"
	"time"
)

// Debug Debug 日志
func (a *AppLoggerImpl) Debug(ctx context.Context, v ...interface{}) {
	a.write(ctx, logtypes.DebugLv, logtypes.DEBUG, v...)
}

// Debugf Debugf 日志
func (a *AppLoggerImpl) Debugf(ctx context.Context, format string, v ...interface{}) {
	a.write(ctx, logtypes.DebugLv, logtypes.DEBUG, fmt.Sprintf(format, v...))
}

// Info Info 日志
func (a *AppLoggerImpl) Info(ctx context.Context, v ...interface{}) {
	a.write(ctx, logtypes.InfoLv, logtypes.INFO, v...)
}

// Infof Infof 日志
func (a *AppLoggerImpl) Infof(ctx context.Context, format string, v ...interface{}) {
	a.write(ctx, logtypes.InfoLv, logtypes.INFO, fmt.Sprintf(format, v...))
}

// Warn Warn 日志
func (a *AppLoggerImpl) Warn(ctx context.Context, v ...interface{}) {
	a.write(ctx, logtypes.WarnLv, logtypes.WARN, v...)
}

// Warnf Warnf 日志
func (a *AppLoggerImpl) Warnf(ctx context.Context, format string, v ...interface{}) {
	a.write(ctx, logtypes.WarnLv, logtypes.WARN, fmt.Sprintf(format, v...))
}

// Error Error 日志
func (a *AppLoggerImpl) Error(ctx context.Context, v ...interface{}) {
	a.write(ctx, logtypes.ErrorLv, logtypes.ERROR, v...)
}

// Errorf Errorf 日志
func (a *AppLoggerImpl) Errorf(ctx context.Context, format string, v ...interface{}) {
	a.write(ctx, logtypes.ErrorLv, logtypes.ERROR, fmt.Sprintf(format, v...))
}

// Fatal fatal日志
func (a *AppLoggerImpl) Fatal(ctx context.Context, v ...interface{}) {
	a.write(ctx, logtypes.FatalLv, logtypes.FATAL, v...)
	os.Exit(1)
}

// Fatalf fatal日志
func (a *AppLoggerImpl) Fatalf(ctx context.Context, format string, v ...interface{}) {
	a.write(ctx, logtypes.FatalLv, logtypes.FATAL, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (a *AppLoggerImpl) write(ctx context.Context, flagX logtypes.Level, flag logtypes.LevelFlag, v ...interface{}) {
	tid, _, err := excontext.GetTraceIDExist(ctx)
	if err != nil && a != nil && a.PrintTraceIdErr {
		MainLogger.Warn(excontext.MainTraceCtx(), err.Error())
	}

	if a == nil {
		// 日期时间|日志级别|线程名/traceId|文件名|记录内容 行结尾符
		logData := logFormat(ctx, flag, 5, tid, fmt.Sprint(v...))
		fmt.Print(logData)
		return
	}
	if a.LogLevel > flagX {
		return
	}
	logger := a.Loggers[flag]

	// 日期时间|日志级别|线程名/traceId|文件名|记录内容 行结尾符
	logData := a.LogFormatFunc(ctx, flag, a.CallerDeep, tid, fmt.Sprint(v...))
	if logger == nil {
		fmt.Print(logData)
		return
	}
	// opt配置里不打印文件名，这里为 calldepth无所谓
	err = logger.Output(0, logData)
	if err != nil {
		fmt.Print(logData)
	}
	return
}

func logFormat(ctx context.Context, flag logtypes.LevelFlag, callDeep int, tid, body string) string {

	return fmt.Sprintf(traceIDFormatStr,
		time.Now().Format("2006-01-02 15:04:05.000"),
		flag,
		tid,
		callerFile(callDeep),
		body)
}

func callerFile(calldepth int) string {
	file, line := exruntime.CallerFile(calldepth)
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return fmt.Sprintf("%s:%d", file, line)
}
