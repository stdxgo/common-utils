package inits

import (
	"context"
	"fmt"
	"github.com/stdxgo/common-utils/exruntime"
	"sort"
	"strings"
	"sync"
)

var (
	initFun0       initFuncArr
	initFun0Locker sync.Mutex
	initDone       bool

	defaultWeight = 1000
)

// RegisterInitFunc 以默认权重注册方法
func RegisterInitFunc(f func(ctx context.Context)) {
	registerInit(defaultWeight, f)
}

// RegisterInitFuncWithWeight 注册带权重的init方法
func RegisterInitFuncWithWeight(f func(ctx context.Context), w int) {
	registerInit(w, f)
}

func registerInit(w int, ctxFunc func(ctx context.Context)) {

	file, line := exruntime.CallerFile(2)
	initFun0Locker.Lock()
	defer initFun0Locker.Unlock()
	initFun0 = append(initFun0, initFunc{
		f:     ctxFunc,
		w:     w,
		paths: filePathErasePrefix(file),
		line:  line,
	})
}

func filePathErasePrefix(file string) []string {
	return strings.Split(file, "/")
}

// only for test
func clearReg() {
	initFun0Locker.Lock()
	defer initFun0Locker.Unlock()
	initFun0 = nil
	initDone = false
}

func runRegisteredInit(ctx context.Context) {

	if initDone {
		return
	}
	sort.Sort(initFun0)
	initRunnerPath()

	defer func() {
		initDone = true
	}()
	for _, if0 := range initFun0 {
		fmt.Printf("执行初始化注册方法(权重:%.5d)：%s\n", if0.w, if0.runnerPath)
		if if0.f != nil {
			if0.f(ctx)
		}
	}
	return
}

// RunRegisteredInit 执行注册的init方法
func RunRegisteredInit(ctx context.Context) {
	initFun0Locker.Lock()
	defer initFun0Locker.Unlock()
	runRegisteredInit(ctx)
}
func initRunnerPath() {
	initFun0Foreach := func(f func(ifc *initFunc)) {
		for i := range initFun0 {
			f(&initFun0[i])
		}
	}
	initFun0Foreach(func(ifc *initFunc) {
		ifc.runnerPath = fmt.Sprintf("%s:%d", strings.Join(ifc.paths, "/"), ifc.line)
	})
}
