package inits

import "context"

type initFunc struct {
	f          func(ctx context.Context)
	w          int
	paths      []string
	line       int
	runnerPath string
}

type initFuncArr []initFunc

func (ifa initFuncArr) Len() int {
	return len(ifa)
}

func (ifa initFuncArr) Less(i, j int) bool {
	if ifa[i].w > ifa[j].w {
		return true
	}
	return false
}

func (ifa initFuncArr) Swap(i, j int) {
	ifa[i], ifa[j] = ifa[j], ifa[i]
}
