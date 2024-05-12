package strutil

import "math/rand"

var (
	stableStrSeed = rand.New(rand.NewSource(1))
	stableNumSeed = rand.New(rand.NewSource(1))
)

// GenStr4Test 随机所有字符（数字、大小写字母）:
// 仅用于测试，稳定生成
func GenStr4Test(le int) string {

	return genStrWithInputsAndRand(stableStrSeed, digitLetterBytes, le)
}

// GenNum4Test 数字 :
// 仅用于测试，稳定生成
func GenNum4Test(le int) string {
	return genStrWithInputsAndRand(stableNumSeed, digitBytes, le)
}

// GenStr 随机所有字符（数字、大小写字母）
func GenStr(le int) string {
	return genStrWithInputs(digitLetterBytes, le)
}

// GenNum 数字
func GenNum(le int) string {
	return genStrWithInputs(digitBytes, le)
}

// GenLetters 大小写字母
func GenLetters(le int) string {
	return genStrWithInputs(letterBytes, le)
}

// GenLowerLettersAndNum 小写字母
func GenLowerLettersAndNum(le int) string {
	return genStrWithInputs(digitLowerLetterBytes, le)
}

// GenByBytes 根据输入生成指定长度字符
func GenByBytes(in []byte, le int) string {
	if len(in) < 1 {
		return ""
	}
	return genStrWithInputs(in, le)
}

// GenByBytesWithRand 根据输入生成指定长度字符
func GenByBytesWithRand(rg *rand.Rand, in []byte, le int) string {
	if len(in) < 1 {
		return ""
	}
	return genStrWithInputsAndRand(rg, in, le)
}
