package strutil

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

const (
	Digit           = "0123456789"
	LowerCaseLetter = "abcdefghijklmnopqrstuvwxyz"
	UpperCaseLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	digitBytes            = []byte(Digit)
	digitLowerLetterBytes = []byte(Digit + LowerCaseLetter)
	letterBytes           = []byte(LowerCaseLetter + UpperCaseLetter)
	digitLetterBytes      = []byte(Digit + LowerCaseLetter + UpperCaseLetter)
)

func genStrWithInputs(bytes []byte, length int) string {

	randomGen := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	return genStrWithInputsAndRand(randomGen, bytes, length)
}

func genStrWithInputsAndRand(rg *rand.Rand, bytes []byte, length int) string {

	bytesLen := len(bytes)
	retVal := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		retVal = append(retVal, bytes[rg.Intn(bytesLen)])
	}
	return string(retVal)
}

// MD5 md5 func
func MD5(data []byte) string {
	tmp := md5.Sum(data)
	return hex.EncodeToString(tmp[:])
}
