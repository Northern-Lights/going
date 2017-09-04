package crypt

import (
	"bytes"
	"testing"
)

func TestAesCbc(t *testing.T) {
	data := []byte("what do ya want for nothing?")
	key := []byte("\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a")
	iv := []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	enc, err := AesEncryptCbc(data, key, iv)
	if err != nil {
		t.Errorf("Error in AesEncryptCbc: %s\n", err.Error())
	}
	dec, err := AesDecryptCbc(enc, key, iv)
	if err != nil {
		t.Errorf("Error in AesDecryptCbc: %s\n", err.Error())
	}

	// TODO: Check for equivalence. Requires removing padding from dec.
	if !bytes.Contains(dec, data) {
		t.Error("Decrypted bytes != plaintext bytes")
	}
}
