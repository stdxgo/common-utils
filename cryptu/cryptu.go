package cryptu

import (
	"errors"
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
