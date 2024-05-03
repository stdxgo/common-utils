package cryptutil

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	testx()
	testx()
	testx()
	testx()
	testx()
	testx()
	testx()
	testx()
	testx()
	testx()
	testx()
}

func testx() {
	key := "12345678901234567890"
	b, e := AesCtrEncrypt("asdasdasdasdasdasd", key)
	if e != nil {
		panic(e)
	}
	fmt.Println(b)

	result, err := AesCtrDecrypt(b, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
