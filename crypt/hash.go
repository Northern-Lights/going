package crypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

// A HashAlgorithm takes a slice of bytes, hashes it with an algorithm, and
// returns the hashed byte slice, which is more flexible/generic than a fixed
// array.
type HashAlgorithm func([]byte) []byte

// Hash provides quick access to hashing functions through a map.
var Hash = map[string](HashAlgorithm){
	"MD5":    Md5,
	"SHA1":   Sha1,
	"SHA256": Sha256,
	"SHA512": Sha512,
}

// Md5 performs MD5 on a byte slice and returns the MD5 hash as a byte slice.
func Md5(dataIn []byte) (dataOut []byte) {
	b := md5.Sum(dataIn)
	dataOut = b[:]
	return
}

// Sha1 performs SHA1 on a byte slice and returns the SHA1 hash as a byte slice.
func Sha1(dataIn []byte) (dataOut []byte) {
	b := sha1.Sum(dataIn)
	dataOut = b[:]
	return
}

// Sha256 performs SHA256 on a byte slice and returns the SHA256 hash as a byte slice.
func Sha256(dataIn []byte) (dataOut []byte) {
	b := sha256.Sum256(dataIn)
	dataOut = b[:]
	return
}

// Sha512 performs SHA512 on a byte slice and returns the SHA512 hash as a byte slice.
func Sha512(dataIn []byte) (dataOut []byte) {
	b := sha512.Sum512(dataIn)
	dataOut = b[:]
	return
}
