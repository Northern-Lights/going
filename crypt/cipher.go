package crypt

// Xor performs single-byte XOR on a data buffer.
func Xor(data *([]byte), key byte) {
	for i := 0; i < len(*data); i++ {
		(*data)[i] ^= key
	}
}
