package ezcrypt

import (
	"bytes"
	"encoding/gob"
)

func EncryptGob(key []byte, object interface{}) (encrypted []byte, nonce []byte, err error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(object)
	content := buffer.Bytes()
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

func DecryptGob(key []byte, object interface{}, encrypted []byte, nonce []byte) (err error) {
	content, err := Decrypt(key, encrypted, nonce)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(content)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(object)
	if err != nil {
		return err
	}
	return nil
}
