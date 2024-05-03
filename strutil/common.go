package strutil

import "strings"

// ReplaceNewLineChar 去除字符串中的 \r \n
func ReplaceNewLineChar(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	return strings.ReplaceAll(s, "\n", "")
}
