package openssl

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"math"
	"math/big"
)

type KeyFormat uint8  // key format
type blockType string // key block type

const (
	PKCS1 KeyFormat = iota // PKCS#1
	PKCS8                  // PKCS#8
)

const (
	privateBlockType blockType = "RSA PRIVATE KEY"
	publicBlockType  blockType = "PUBLIC KEY"
)

var KeyFormatErr = errors.New("the key format must be PKCS#1 or PKCS#8")
var PrivateKeyErr = errors.New("private key error")
var PublicKeyErr = errors.New("public key error")

var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

// RSAGenerateKey generate public and private keys in pem format
// bits: 512、1024、2048、3072、4096
// format: PKCS1、PKCS8
func RSAGenerateKey(bits int, format KeyFormat) (privateKey string, publicKey string, err error) {
	// private key
	pk, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}

	// key format
	var priKey, pubKey []byte
	switch format {
	case PKCS1:
		priKey, pubKey, err = generatePKCS1Key(pk)
	case PKCS8:
		priKey, pubKey, err = generatePKCS8Key(pk)
	default:
		return "", "", KeyFormatErr
	}
	if err != nil {
		return
	}

	// pem
	privateKey, err = encodePEMPrivateKey(priKey)
	if err != nil {
		return
	}
	publicKey, err = encodePEMPublicKey(pubKey)
	return
}

func RSAEncryptByPrivateKey(key string, data []byte) (string, error) {
	pk, err := decodePEMPrivateKey(key)
	if err != nil {
		return "", err
	}

	if len(data)+11 <= pk.Size() {
		b, err := rsa.SignPKCS1v15(rand.Reader, pk, 0, data)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(b), nil
	}

	bs := bytesSplit(data, pk.Size()-11)
	ciphertext := make([]byte, 0, len(bs)*pk.Size())
	for i, _ := range bs {
		if tmp, err := rsa.SignPKCS1v15(rand.Reader, pk, 0, bs[i]); err != nil {
			return "", err
		} else {
			ciphertext = append(ciphertext, tmp...)
		}
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func RSAEncryptByPublicKey(key string, data []byte) (string, error) {
	pk, err := decodePEMPublicKey(key)
	if err != nil {
		return "", err
	}

	if len(data)+11 <= pk.Size() {
		b, err := rsa.EncryptPKCS1v15(rand.Reader, pk, data)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(b), nil
	}

	bs := bytesSplit(data, pk.Size()-11)
	ciphertext := make([]byte, 0, len(bs)*pk.Size())
	for i, _ := range bs {
		if tmp, err := rsa.EncryptPKCS1v15(rand.Reader, pk, bs[i]); err != nil {
			return "", err
		} else {
			ciphertext = append(ciphertext, tmp...)
		}
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func RSADecryptByPrivateKey(key string, data []byte) ([]byte, error) {
	pk, err := decodePEMPrivateKey(key)
	if err != nil {
		return nil, err
	}

	if len(data) <= pk.Size() {
		return rsa.DecryptPKCS1v15(rand.Reader, pk, data)
	}

	bs := bytesSplit(data, pk.Size())
	plaintext := make([]byte, 0, len(bs)*pk.Size())
	for i, _ := range bs {
		if tmp, err := rsa.DecryptPKCS1v15(rand.Reader, pk, bs[i]); err != nil {
			return nil, err
		} else {
			plaintext = append(plaintext, tmp...)
		}
	}

	return plaintext, nil
}

func RSADecryptByPublicKey(key string, data []byte) ([]byte, error) {
	pk, err := decodePEMPublicKey(key)
	if err != nil {
		return nil, err
	}

	if len(data) <= pk.Size() {
		return publicKeyDecrypt(pk, 0, nil, data)
	}

	bs := bytesSplit(data, pk.Size())
	plaintext := make([]byte, 0, len(bs)*pk.Size())
	for i, _ := range bs {
		if tmp, err := publicKeyDecrypt(pk, 0, nil, bs[i]); err != nil {
			return nil, err
		} else {
			plaintext = append(plaintext, tmp...)
		}
	}

	return plaintext, nil
}

// generatePKCS1Key generate PKCS#1 key from private key
func generatePKCS1Key(pk *rsa.PrivateKey) (privateKey []byte, publicKey []byte, err error) {
	privateKey = x509.MarshalPKCS1PrivateKey(pk)
	publicKey, err = x509.MarshalPKIXPublicKey(pk.Public())
	return
}

// generatePKCS8Key generate PKCS#8 key from private key
func generatePKCS8Key(pk *rsa.PrivateKey) (privateKey []byte, publicKey []byte, err error) {
	privateKey, err = x509.MarshalPKCS8PrivateKey(pk)
	if err != nil {
		return
	}
	publicKey, err = x509.MarshalPKIXPublicKey(pk.Public())
	return
}

// encodePEMKey encode private/public key to pem
func encodePEMKey(key []byte, bt blockType) (string, error) {
	buf := bytes.NewBuffer(nil)
	block := &pem.Block{
		Type:  string(bt),
		Bytes: key,
	}
	if err := pem.Encode(buf, block); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// encodePEMPrivateKey encode private key to pem
func encodePEMPrivateKey(key []byte) (string, error) {
	return encodePEMKey(key, privateBlockType)
}

// encodePEMPublicKey encode public key to pem
func encodePEMPublicKey(key []byte) (string, error) {
	return encodePEMKey(key, publicBlockType)
}

// decodePEMPrivateKey decode pem to private key
func decodePEMPrivateKey(key string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, PrivateKeyErr
	}

	if privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return privateKey, nil
	}
	pkAny, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey, ok := pkAny.(*rsa.PrivateKey)
	if !ok {
		return nil, PrivateKeyErr
	}
	return privateKey, nil
}

// decodePEMPublicKey decode pem to public key
func decodePEMPublicKey(key string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, PublicKeyErr
	}

	pkAny, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := pkAny.(*rsa.PublicKey)
	if !ok {
		return nil, PublicKeyErr
	}
	return publicKey, nil
}

// bytesSplit split the data according to the specified length
func bytesSplit(data []byte, partLen int) (result [][]byte) {
	dataLen := len(data)
	if dataLen <= partLen || partLen <= 0 {
		result = append(result, data)
		return
	}
	// split
	times := int(math.Ceil(float64(dataLen) / float64(partLen)))
	var left, right int
	for i := 0; i < times; i++ {
		right = (i + 1) * partLen
		if right > dataLen {
			right = dataLen
		}
		result = append(result, data[left:right])
		left += partLen
	}
	return
}

// pkcs1v15HashInfo from rsa.pkcs1v15HashInfo
func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	// Special case: crypto.Hash(0) is used to indicate that the data is
	// signed directly.
	if hash == 0 {
		return inLen, nil, nil
	}

	hashLen = hash.Size()
	if inLen != hashLen {
		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
	}
	prefix, ok := hashPrefixes[hash]
	if !ok {
		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
	}
	return
}

// encrypt from rsa.encrypt
func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// padding
func padding(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

// unPadding
func unPadding(input []byte) (out []byte) {
	n := len(input)
	t := 2
	for i := 2; i < n; i++ {
		if input[i] == 0xff {
			t = t + 1
		} else {
			if input[i] == input[0] {
				t = t + int(input[1])
			}
			break
		}
	}
	out = make([]byte, n-t)
	copy(out, input[t:])
	return
}

// publicKeyDecrypt
func publicKeyDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sign []byte) (out []byte, err error) {
	hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := (pub.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, rsa.ErrMessageTooLong
	}

	c := new(big.Int).SetBytes(sign)
	m := encrypt(&big.Int{}, pub, c)
	em := padding(m.Bytes(), k)
	out = unPadding(em)

	err = nil
	return
}
