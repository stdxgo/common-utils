package internal

import (
	"errors"
	"fmt"
	"github.com/stdxgo/common-utils/logging/logtypes"
	"strings"
)

var (
	lvFlagMap2Lv = map[logtypes.LevelFlag]logtypes.Level{
		logtypes.DEBUG: logtypes.DebugLv,
		logtypes.INFO:  logtypes.InfoLv,
		logtypes.WARN:  logtypes.WarnLv,
		logtypes.ERROR: logtypes.ErrorLv,
		logtypes.FATAL: logtypes.FatalLv,
	}
	lvMap2LvFlag = map[logtypes.Level]logtypes.LevelFlag{}
)

func init() {
	for flag, level := range lvFlagMap2Lv {
		lvMap2LvFlag[level] = flag
	}
}

type logFileType string

const (
	logFileTypeSystem logFileType = "system"
	logFileTypeError  logFileType = "error"
)

// LogLevel 日志级别转为int表示方式
func LogLevel(llv string) logtypes.Level {
	if lv, ok := lvFlagMap2Lv[logtypes.LevelFlag(strings.ToUpper(llv))]; ok {
		return lv
	}
	fmt.Printf("\nWARNING: 日志级别配置(%s)未存在于系统内置，按默认日志级别INFO输出 \n\n", llv)
	return logtypes.InfoLv
}

// ValidLogFileType valid logFileType
func ValidLogFileType(lft logFileType) error {
	if !validLFT[lft] {
		return errors.New("logFileType 不合法，支持类型：system/error")
	}
	return nil
}

var (
	levelFlags   = []logtypes.LevelFlag{logtypes.DEBUG, logtypes.INFO, logtypes.WARN, logtypes.ERROR, logtypes.FATAL}
	logFileTypes = []logFileType{logFileTypeSystem, logFileTypeSystem, logFileTypeSystem, logFileTypeError, logFileTypeError}

	validLFT = map[logFileType]bool{
		logFileTypeSystem: true,
		logFileTypeError:  true,
	}
)
