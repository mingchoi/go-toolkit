package ezcrypt_test

import (
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/mingchoi/toolkit/ezcrypt"
)

var path string
var exampleKey string = "qwertyuiopasdfghjklzxcvbnm123456"

func TestEncryptDecryptString(t *testing.T) {

	if runtime.GOOS == "windows" {
		path = filepath.FromSlash("./encryptedData")
	} else {
		path = "/tmp/encryptedData"
	}

	confidentialMessage := "This message is confidential, only restricted person have premission to read."

	// Encryption
	t.Log("\n=== Encryption ===\n")
	t.Log("Method: AES-256 GCM\n")
	content, iv, err := ezcrypt.EncryptString(exampleKey, confidentialMessage)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Nonce: %x\nEncrypted Content: \n%s\n", iv, hex.Dump(content))

	// Write to file
	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		t.Error(err)
	}

	// Read from file
	cipherData, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}

	// Decryption
	t.Log("=== Decryption ===\n")
	decryptedMessage, err := ezcrypt.DecryptString(exampleKey, cipherData, iv)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s\n\n", decryptedMessage)

	if decryptedMessage != confidentialMessage {
		t.Error("Decrypted message incorrect")
	}
}
