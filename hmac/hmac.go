package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func SHA1(secret []byte, data []byte) []byte {
	h := hmac.New(sha1.New, secret)
	h.Write(data)
	return h.Sum(nil)
}

func SHA1ToString(secret []byte, data []byte) string {
	return hex.EncodeToString(SHA1(secret, data))
}

func SHA224(secret []byte, data []byte) []byte {
	h := hmac.New(sha256.New224, secret)
	h.Write(data)
	return h.Sum(nil)
}

func SHA224ToString(secret []byte, data []byte) string {
	return hex.EncodeToString(SHA224(secret, data))
}

func SHA256(secret []byte, data []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	return h.Sum(nil)
}

func SHA256ToString(secret []byte, data []byte) string {
	return hex.EncodeToString(SHA256(secret, data))
}

func SHA384(secret []byte, data []byte) []byte {
	h := hmac.New(sha512.New384, secret)
	h.Write(data)
	return h.Sum(nil)
}

func SHA384ToString(secret []byte, data []byte) string {
	return hex.EncodeToString(SHA384(secret, data))
}

func SHA512(secret []byte, data []byte) []byte {
	h := hmac.New(sha512.New, secret)
	h.Write(data)
	return h.Sum(nil)
}

func SHA512ToString(secret []byte, data []byte) string {
	return hex.EncodeToString(SHA512(secret, data))
}
