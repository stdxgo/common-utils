package cryptutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/stdxgo/common-utils/strutil"
)

// AesCtrEncrypt do encrypt
func AesCtrEncrypt(plainText, key string) (string, error) {

	newPlainBytes, err := mixInput(plainText)
	if err != nil {
		return "", errors.New("encrypt fail:" + err.Error())
	}
	return aesCtrEncryptRaw(newPlainBytes, []byte(key))
}

// AesCtrDecrypt do decrypt
func AesCtrDecrypt(encryptData, key string) (string, error) {

	result, err := aesCtrDecryptRaw(encryptData, key)
	if err != nil {
		return "", err
	}
	ei, err := dividePlainText(result)
	if err != nil {
		return "", errors.New("decrypt fail:" + err.Error())
	}
	return ei.V, nil
}

func aesCtrEncryptRaw(plainText, key []byte) (string, error) {

	iv := strutil.GenRandomDigitLowerLetter(4)
	dst, err := aesCtrEncrypt(plainText, key, bytes.Repeat([]byte(iv), 4))
	if err != nil {
		return "", err
	}
	result := Base64Encoding(dst)
	return iv + string(result), nil
}

func aesCtrDecryptRaw(encryptData, key string) ([]byte, error) {

	iv := bytes.Repeat([]byte(encryptData[:4]), 4)
	encryptData = encryptData[4:]
	inp, err := Base64Decoding([]byte(encryptData))
	if err != nil {
		return []byte{}, err
	}
	return aesCtrEncrypt(inp, []byte(key), iv)
}

func aesCtrEncrypt(plainText, key, iv []byte) ([]byte, error) {
	key = []byte(strutil.MD5(key))
	// 第一步：创建aes密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//第二步：创建分组模式ctr ,
	stream := cipher.NewCTR(block, iv)
	//第三步：加密
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)
	return dst, nil
}

var (
	urle = base64.URLEncoding
)

// Base64Encoding base64 encode
func Base64Encoding(en []byte) []byte {
	result := make([]byte, urle.EncodedLen(len(en)))
	urle.Encode(result, en)
	return result
}

// Base64Decoding base64 decode
func Base64Decoding(de []byte) ([]byte, error) {
	result := make([]byte, urle.DecodedLen(len(de)))
	le, err := urle.Decode(result, de)
	if err != nil {
		return nil, err
	}
	return result[:le], nil
}
