package excontext

import (
	"context"
	"fmt"
	"github.com/stdxgo/common-utils/strutil"
	"runtime"
)

type ctxKey string
type ctxVal string

const (
	// TraceCtxKey trace context key
	TraceCtxKey ctxKey = "Trace-Access-Id"
)

func randomTraceValue() ctxVal {
	return ctxVal(strutil.GenRandomDigitLowerLetter(16))
}

// NewExTraceContext 创建一个ExTraceContext
func NewExTraceContext(ctx context.Context) context.Context {
	return trace(NewExContext(ctx), randomTraceValue())
}

// NewTraceContext 创建一个 TraceContext
func NewTraceContext(ctx context.Context) context.Context {
	return trace(ctx, randomTraceValue())
}

func NewTraceContextWithId(ctx context.Context, traceId string) context.Context {
	return trace(ctx, ctxVal(traceId))
}

func trace(ctx context.Context, val ctxVal) context.Context {
	return context.WithValue(ctx, TraceCtxKey, val)
}

var (
	mainTraceCtx = trace(context.Background(), "main")
)

// MainTraceCtx 获取主TraceContext
func MainTraceCtx() context.Context {
	return mainTraceCtx
}

// EmptyTraceCtx 获取空的TraceContext
func EmptyTraceCtx() context.Context {
	return trace(context.TODO(), "")
}

// GetTraceID 从Context中获取TraceID
func GetTraceID(ctx context.Context) string {
	tid, _, _ := getTraceIDExist(ctx)
	return tid
}

// GetTraceIDExist 从context中获取TraceID与其存在性
func GetTraceIDExist(ctx context.Context) (string, bool, error) {
	return getTraceIDExist(ctx)
}

func getTraceIDExist(ctx context.Context) (string, bool, error) {
	v := ctx.Value(TraceCtxKey)
	var err error
	if v == nil {
		var stackBytes [4096]byte // 4KB
		le := runtime.Stack(stackBytes[:], false)
		err = fmt.Errorf("获取traceId时,val值为nil :\n%s\n", stackBytes[:le])
		return "-", false, err
	}
	vs, ok := v.(ctxVal)
	if !ok {
		var stackBytes [4096]byte // 4KB
		le := runtime.Stack(stackBytes[:], false)
		err = fmt.Errorf("获取traceId时,val值类型不为ctxVal :\n%s\n", stackBytes[:le])
		return "-", false, err
	}
	return string(vs), true, nil
}
