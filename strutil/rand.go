package strutil

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

const (
	digit           = "0123456789"
	lowerCaseLetter = "abcdefghijklmnopqrstuvwxyz"
	upperCaseLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	strLetter        = []byte(digit + lowerCaseLetter + upperCaseLetter)
	digitLowerLetter = []byte(digit + lowerCaseLetter)
)

// GenRandomString 生成随机字符串
func GenRandomString(length int) string {
	return genStrByLenAndBaseStr(strLetter, length)
}

// GenRandomDigitLowerLetter 生成小写字母与数字随机串
func GenRandomDigitLowerLetter(length int) string {
	return genStrByLenAndBaseStr(digitLowerLetter, length)
}

func genStrByLenAndBaseStr(bytes []byte, length int) string {

	bytesLen := len(bytes)
	retVal := make([]byte, 0, length)

	randomGen := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	for i := 0; i < length; i++ {
		retVal = append(retVal, bytes[randomGen.Intn(bytesLen)])
	}
	return string(retVal)
}

// MD5 md5 func
func MD5(data []byte) string {
	tmp := md5.Sum(data)
	return hex.EncodeToString(tmp[:])
}
