package ezcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPWToByte(content []byte, salt []byte) []byte {
	sum, err := bcrypt.GenerateFromPassword(append(content, salt...), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return sum[:]
}

func HashPW(content []byte, salt []byte) string {
	return string(HashPWToByte(content, salt))
}

func ValidPW(content []byte, salt []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, append(content, salt...))
	return err == nil
}
