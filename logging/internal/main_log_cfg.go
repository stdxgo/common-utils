package internal

import "github.com/stdxgo/common-utils/logging/logtypes"

// SetMainLogLevel 设置mainLog的日志级别
func SetMainLogLevel(level logtypes.Level) {
	MainLoggerOprWithLock(func(mLogger *AppLoggerImpl) {
		mLogger.LogLevel = level
	})
}

// SetMainLogCallerDeep 设置mainLog的callerDeep
func SetMainLogCallerDeep(deep int) {
	MainLoggerOprWithLock(func(mLogger *AppLoggerImpl) {
		mLogger.CallerDeep = callerDeepWrap(deep)
	})
}
