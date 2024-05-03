package excontext

import (
	"context"
	"time"
)

// NewExContext 创建一个ExContext
func NewExContext(ctx context.Context) context.Context {
	return &newKeepKVContext{
		oldCtx: ctx,
		newCtx: context.Background(),
	}
}

type newKeepKVContext struct {
	oldCtx context.Context
	newCtx context.Context
}

func (a *newKeepKVContext) Deadline() (deadline time.Time, ok bool) {
	return a.newCtx.Deadline()
}

func (a *newKeepKVContext) Done() <-chan struct{} {
	return a.newCtx.Done()
}

func (a *newKeepKVContext) Err() error {
	return a.newCtx.Err()
}

func (a *newKeepKVContext) Value(key interface{}) interface{} {
	return a.oldCtx.Value(key)
}
