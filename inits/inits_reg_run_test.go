package inits

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRunRegisteredInit(t *testing.T) {
	RegisterInitFunc(func(ctx context.Context) {})
	RegisterInitFunc(func(ctx context.Context) {})
	RegisterInitFunc(func(ctx context.Context) {})
	RegisterInitFuncWithWeight(func(ctx context.Context) {}, 0)
	RegisterInitFuncWithWeight(func(ctx context.Context) {}, 1001)
	wg := sync.WaitGroup{}
	wg.Add(2)
	RegisterInitFunc(func(ctx context.Context) {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("done")
					return
				default:
					fmt.Println("sleep---")
					time.Sleep(time.Second)
				}
			}
		}()
	})

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)
		cancel()
	}()

	RunRegisteredInit(ctx)
	RunRegisteredInit(ctx)

	clearReg()
	wg.Wait()
}
