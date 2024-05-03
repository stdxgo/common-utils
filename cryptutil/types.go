package cryptutil

import (
	"encoding/json"
	"math/rand"
)

type encryptInfo struct {
	X int    `json:"x"` //
	V string `json:"v"` // value
}

func mixInput(plainText string) ([]byte, error) {
	return json.Marshal(encryptInfo{
		V: plainText,
		X: rand.Int() % 1000,
	})
}

func dividePlainText(mixedInput []byte) (encryptInfo, error) {

	var ei encryptInfo
	err := json.Unmarshal(mixedInput, &ei)
	return ei, err
}
