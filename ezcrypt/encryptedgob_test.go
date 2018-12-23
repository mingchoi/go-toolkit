package ezcrypt_test

import (
	"testing"

	"github.com/mingchoi/toolkit/ezcrypt"
)

type Gender bool

const (
	M Gender = true
	F Gender = false
)

type Person struct {
	Name   string
	Age    int
	Gender Gender
}

func TestEncryptGob(t *testing.T) {
	// Encrypt
	peter := Person{"Peter", 20, M}
	encrypted, nonce, err := ezcrypt.EncryptGob([]byte(exampleKey), &peter)
	if err != nil {
		t.Error(err)
	}

	// Decrypt
	person := Person{}
	err = ezcrypt.DecryptGob([]byte(exampleKey), &person, encrypted, nonce)
	if err != nil {
		t.Error(err)
	}

	if person != peter {
		t.Error("Decrypted object incorrect")
	}

}
