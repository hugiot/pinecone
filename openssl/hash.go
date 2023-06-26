package openssl

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func MD5(p []byte) []byte {
	h := md5.New()
	h.Write(p)
	return h.Sum(nil)
}

func MD5ToString(p []byte) string {
	return hex.EncodeToString(MD5(p))
}

func SHA1(p []byte) []byte {
	h := sha1.New()
	h.Write(p)
	return h.Sum(nil)
}

func SHA1ToString(p []byte) string {
	return hex.EncodeToString(SHA1(p))
}

func SHA256(p []byte) []byte {
	h := sha256.New()
	h.Write(p)
	return h.Sum(nil)
}

func SHA256ToString(p []byte) string {
	return hex.EncodeToString(SHA256(p))
}
