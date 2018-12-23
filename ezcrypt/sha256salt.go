package ezcrypt

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func HashToByte(content []byte, salt []byte) []byte {
	sum := sha256.Sum256(append(content, salt...))
	return sum[:]
}

func HashWithoutSymbol(content []byte, salt []byte) string {
	str := (string)(content)
	str = strings.Replace(str, "-", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, " ", "", -1)
	return hex.EncodeToString(HashToByte([]byte(str), salt))
}

func HashOriginal(content []byte, salt []byte) string {
	return hex.EncodeToString(HashToByte(content, salt))
}
