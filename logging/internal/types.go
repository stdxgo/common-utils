package internal

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/stdxgo/common-utils/logging/logtypes"
	"log"
	"os"
	"strings"
	"time"
)

const (
	// KB kb
	KB int64 = 1024
	// MB mb
	MB = 1024 * KB
)

// AppLoggerImpl AppLog实现
type AppLoggerImpl struct {
	CallerDeep      int
	LogLevel        logtypes.Level
	Loggers         map[logtypes.LevelFlag]*log.Logger
	LogFormatFunc   logtypes.LogFormatFunc
	PrintTraceIdErr bool
}

// LogOpt log配置
type LogOpt struct {
	CallerDeep    int
	LogPath       string         `valid:"required~logPath不能为空"`
	LogAgeDay     time.Duration  `valid:"required~logAge不能为空"`
	LogLevel      logtypes.Level `valid:""`
	AppName       string         `valid:"required~Log日志初始化参数AppName不能为空"`
	LogFormatFunc logtypes.LogFormatFunc
}

// GetAppLogger 新建一个AppLogger
func GetAppLogger(lo LogOpt) logtypes.AppLogger {

	if _, err := govalidator.ValidateStruct(&lo); err != nil {
		panic(err.Error())
	}

	var aLoger = AppLoggerImpl{
		Loggers:  make(map[logtypes.LevelFlag]*log.Logger),
		LogLevel: lo.LogLevel,
	}
	logFilePattern := func(ft logFileType, lf logtypes.LevelFlag) string {
		return fmt.Sprintf("%s/%s.%s.%s.%s.log", lo.LogPath, ft, lo.AppName, "%Y%m%d", lf)
	}
	for i := int(lo.LogLevel); i < len(levelFlags); i++ {
		logf, err := rotatelogs.New(
			logFilePattern(logFileTypes[i], levelFlags[i]), rotateOpts(lo.LogAgeDay)...)

		if err != nil {
			panic(fmt.Sprintf("Could not init logging err: %v", err))
		}
		aLoger.Loggers[levelFlags[i]] = log.New(logf, "", 0)
	}
	aLoger.CallerDeep = callerDeepWrap(lo.CallerDeep)
	aLoger.LogFormatFunc = lo.LogFormatFunc
	if aLoger.LogFormatFunc == nil {
		aLoger.LogFormatFunc = logFormat
	}

	aLoger.PrintTraceIdErr = strings.ToLower(os.Getenv("NoteTraceIdErr")) == "true"
	return &aLoger
}

func callerDeepWrap(deep int) int {
	return deep + 4
}

func rotateOpts(logAgeDay time.Duration) []rotatelogs.Option {
	var opts []rotatelogs.Option
	opts = append(opts, rotatelogs.WithMaxAge(logAgeDay*time.Hour*24))
	opts = append(opts, rotatelogs.WithRotationTime(time.Hour*24))
	opts = append(opts, rotatelogs.WithRotationSize(logRotateSize))
	return opts
}

var (
	// 日期时间|日志级别|线程名/traceId|文件名|记录内容 行结尾符
	logFormatStr     = "%s|%s|%s|%s|%s\n"
	traceIDFormatStr = "%s|%s|traceId:%s|%s|%s\n"
	logRotateSize    = 500 * MB
)
