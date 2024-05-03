package exruntime

import (
	"runtime"
)

func CallerFile(calldepth int) (file string, line int) {
	_, file, line, ok := runtime.Caller(calldepth + 1)
	if !ok {
		file = "???"
		line = 0
	}
	return
}
