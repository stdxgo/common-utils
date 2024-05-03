package logging

import (
	"context"
	"github.com/stdxgo/common-utils/inits"
	"github.com/stdxgo/common-utils/logging/internal"
	"github.com/stdxgo/common-utils/logging/logtypes"
	"reflect"
)

func init() {
	inits.RegisterInitFuncWithWeight(initLogFunc, -1)
}

func initLogFunc(ctx context.Context) {
	internal.MainLoggerInitWithLock(func() {
		if internal.MainLogger != nil && !reflect.ValueOf(internal.MainLogger).IsZero() {
			return
		}
		logCfg := logtypes.LogConfig{}
		logCfg.LogLevel = "info"
		internal.DefaultLogCfg(&logCfg)
		// 初始化主日志输出
		internal.InitMainLogger(logCfg)
	})
}
