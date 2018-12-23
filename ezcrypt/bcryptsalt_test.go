package ezcrypt_test

import (
	"testing"

	"github.com/mingchoi/toolkit/ezcrypt"
)

func TestHashPW(t *testing.T) {
	password := []byte("my-stupid-pw")
	salt := []byte(":myName@ezcrypttest")

	hash := ezcrypt.HashPWToByte(password, salt)
	if !ezcrypt.ValidPW(password, salt, hash) {
		t.Error("Valid password unsuccessful")
	}
}
