package strutil

// LastNThIdx 找倒数第N个字符
func LastNThIdx(s string, nth int, b byte) int {
	if nth <= 0 {
		return -1
	}
	cnt := 0
	bs := []byte(s)
	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == b {
			cnt++
		}
		if cnt == nth {
			return i
		}
	}
	return -1
}
