package internal

import (
	"github.com/stdxgo/common-utils/logging/logtypes"
	"sync"
)

// InitMainLogger 初始化mainLog
func InitMainLogger(logCfg logtypes.LogConfig) {
	MainLogger = GetAppLogger(
		LogOpt{
			CallerDeep: 1,
			LogPath:    logCfg.LogPath,
			LogAgeDay:  logCfg.LogAge,
			LogLevel:   LogLevel(logCfg.LogLevel),
			AppName:    logCfg.AppName,
		},
	)
}

var (
	// MainLogger mainLogger
	MainLogger      logtypes.AppLogger = (*AppLoggerImpl)(nil)
	mainLoggerMutex sync.Mutex
)

// DefaultLogCfg logConfig设置默认值
func DefaultLogCfg(logCfg *logtypes.LogConfig) {
	if logCfg.LogPath == "" {
		logCfg.LogPath = "log"
	}
	//
	if logCfg.LogAge <= 3 {
		logCfg.LogAge = 3
	}
	if logCfg.AppName == "" {
		logCfg.AppName = "app"
	}
}

// MainLoggerOprWithLock 修改mainLog配置
func MainLoggerOprWithLock(f func(mLogger *AppLoggerImpl)) {
	if f == nil {
		return
	}
	mainLoggerMutex.Lock()
	defer mainLoggerMutex.Unlock()
	f(MainLogger.(*AppLoggerImpl))
}

// MainLoggerInitWithLock 初始化mainLog
func MainLoggerInitWithLock(f func()) {
	if f == nil {
		return
	}
	mainLoggerMutex.Lock()
	defer mainLoggerMutex.Unlock()
	f()
}
