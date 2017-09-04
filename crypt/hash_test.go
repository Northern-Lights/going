package crypt

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var (
	testString         = []byte("The quick brown fox jumps over the lazy dog")
	md5Reference, _    = hex.DecodeString("9e107d9d372bb6826bd81d3542a419d6")
	sha1Reference, _   = hex.DecodeString("2fd4e1c67a2d28fced849ee1bb76e7391b93eb12")
	sha256Reference, _ = hex.DecodeString("d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592")
	sha512Reference, _ = hex.DecodeString("07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6")
)

func TestHashMd5(t *testing.T) {
	algName := "MD5"
	hash := Hash[algName](testString)
	if !bytes.Equal(md5Reference, hash) {
		t.Errorf("%s expected %x; got %x\n", algName, md5Reference, hash)
	}
}

func TestHashSha1(t *testing.T) {
	algName := "SHA1"
	hash := Hash[algName](testString)
	if !bytes.Equal(sha1Reference, hash) {
		t.Errorf("%s expected %x; got %x\n", algName, md5Reference, hash)
	}
}

func TestHashSha256(t *testing.T) {
	algName := "SHA256"
	hash := Hash[algName](testString)
	if !bytes.Equal(sha256Reference, hash) {
		t.Errorf("%s expected %x; got %x\n", algName, md5Reference, hash)
	}
}

func TestHashSha512(t *testing.T) {
	algName := "SHA512"
	hash := Hash[algName](testString)
	if !bytes.Equal(sha512Reference, hash) {
		t.Errorf("%s expected %x; got %x\n", algName, md5Reference, hash)
	}
}
