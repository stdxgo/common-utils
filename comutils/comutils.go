package comutils

func IF[T any](x bool, v1, v2 T) T {
	if x {
		return v1
	}
	return v2
}

func DoIfTrue(x bool, f func()) {
	if x && f != nil {
		f()
	}
}
