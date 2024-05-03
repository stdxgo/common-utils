package logcfg

import (
	"github.com/stdxgo/common-utils/logging/internal"
	"github.com/stdxgo/common-utils/logging/logtypes"
	"sync"
)

var (
	mainLogInitOnce sync.Once
)

// NewLogging 新建一个AppLogger
func NewLogging(logCfg logtypes.LogConfig) logtypes.AppLogger {
	internal.DefaultLogCfg(&logCfg)
	return internal.GetAppLogger(internal.LogOpt{
		CallerDeep:    1,
		LogPath:       logCfg.LogPath,
		LogAgeDay:     logCfg.LogAge,
		LogLevel:      internal.LogLevel(logCfg.LogLevel),
		AppName:       logCfg.AppName,
		LogFormatFunc: nil,
	})
}

// InitMainLogging 初始化主 logging
func InitMainLogging(logCfg logtypes.LogConfig) {
	mainLogInitOnce.Do(func() {
		internal.MainLoggerInitWithLock(func() {
			internal.DefaultLogCfg(&logCfg)
			internal.InitMainLogger(logCfg)
		})
	})
}
