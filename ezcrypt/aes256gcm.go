package ezcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func GenerateNonce() (nonce []byte, err error) {
	nonce = make([]byte, 12)
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return
}

func Encrypt(key []byte, content []byte, nonce []byte) (encrypted []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	encrypted = aesgcm.Seal(nil, nonce, content, nil)
	return
}

func Decrypt(key []byte, encrypted []byte, nonce []byte) (content []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	content, err = aesgcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return nil, err
	}
	return
}

func EncryptString(keyStr string, message string) (encrypted []byte, nonce []byte, err error) {
	key := []byte(keyStr)
	content := []byte(message)
	nonce, err = GenerateNonce()
	if err != nil {
		return nil, nil, err
	}
	encrypted, err = Encrypt(key, content, nonce)
	if err != nil {
		return nil, nil, err
	}
	return
}

func DecryptString(keyStr string, encrypted []byte, nonce []byte) (content string, err error) {
	key := []byte(keyStr)
	decrypted, err := Decrypt(key, encrypted, nonce)
	if err != nil {
		return "", err
	}
	content = string(decrypted)
	return
}
