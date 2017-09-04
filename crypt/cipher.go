package crypt

import "crypto/aes"
import "crypto/cipher"

// Xor performs single-byte XOR on a data buffer.
func Xor(data *([]byte), key byte) {
	for i := 0; i < len(*data); i++ {
		(*data)[i] ^= key
	}
}

// AesEncryptCbc encrypts a plaintext byte slice in CBC mode.
// It returns an encrypted byte slice.
func AesEncryptCbc(plaintext, key, iv []byte) (enc []byte, err error) {
	const blocksize int = 128 / 8
	numTrailingBytes := len(plaintext) % blocksize
	if numTrailingBytes != 0 {
		paddingSize := blocksize - numTrailingBytes
		for i := 0; i < paddingSize; i++ {
			plaintext = append(plaintext, byte(paddingSize))
		}
	}
	if a, err := aes.NewCipher(key); err == nil {
		enc = make([]byte, len(plaintext))
		c := cipher.NewCBCEncrypter(a, iv)
		c.CryptBlocks(enc, plaintext)
	}
	return
}

// AesDecryptCbc decrypts ciphertext with a key and an IV.
func AesDecryptCbc(ciphertext, key, iv []byte) (dec []byte, err error) {
	const blocksize byte = 128 / 8
	a, err := aes.NewCipher(key)
	if err == nil {
		dec = make([]byte, len(ciphertext))
		c := cipher.NewCBCDecrypter(a, iv)
		c.CryptBlocks(dec, ciphertext)
	} else {
		return
	}
	// FIXME: Use PKCS7 padding, but allow passthrough if non-conformant
	// Now remove padding (assuming PKCS5 padding)
	if paddingByte := dec[len(dec)-1]; paddingByte < blocksize {
		// All of the padding bytes should be the same
		paddingBytes := dec[len(dec)-1-int(paddingByte):]
		for i := byte(0); i < paddingByte; i++ {
			// If not, we don't know what to do. Just return the whole block.
			if paddingBytes[i] != paddingByte {
				return
			}
		}
		dec = dec[:len(dec)-1-int(paddingByte)]
	}
	return
}
