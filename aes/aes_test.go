package aes_test

import (
	"fmt"
	"testing"

	"github.com/go-iot-platform/aes"
)

func TestContainer(t *testing.T) {
	key := "1234567891234567"
	message := "test_msg"

	result := testEncrypt(t, key, message)
	testDecrypt(t, key, result)

}
func testEncrypt(t *testing.T, key string, message string) string {
	result, err := aes.Encrypt(key, message)
	if err != nil || result == "" {
		fmt.Printf("err:%+v\n", err)
		t.Fatal("encrypt fail")
	}
	return result
}
func testDecrypt(t *testing.T, key string, message string) {
	result, err := aes.Decrypt(key, message)
	if err != nil || result == "" {
		fmt.Printf("err:%+v\n", err)
		t.Fatal("decrypt fail")
	}
}
