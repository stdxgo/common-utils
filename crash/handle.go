package crash

import (
	"context"
	"fmt"
	"github.com/stdxgo/common-utils/excontext"
	"runtime/debug"
)

// Handler call f when crash, must call it in defer
func Handler(ctx context.Context, f func()) {

	if re := recover(); re != nil {
		fmt.Println(excontext.GetTraceID(ctx), string(debug.Stack()))
		if f != nil {
			f()
		}
	}
}

// HandlerStk call f when crash, must call it in defer
func HandlerStk(ctx context.Context, f func(stk []byte)) {
	if re := recover(); re != nil {
		if f != nil {
			f(debug.Stack())
		}
	}
}
