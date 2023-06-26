package openssl

import "testing"

func TestMD5ToString(t *testing.T) {
	if s := MD5ToString([]byte("hello world")); s != "5eb63bbbe01eeed093cb22bb8f5acdc3" {
		t.Errorf("MD5ToString(hello world) want return 5eb63bbbe01eeed093cb22bb8f5acdc3, but return %s", s)
	}
}

func TestSHA1ToString(t *testing.T) {
	if s := SHA1ToString([]byte("hello world")); s != "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed" {
		t.Errorf("SHA1ToString(hello world) want return 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed, but return %s", s)
	}
}

func TestSHA256ToString(t *testing.T) {
	if s := SHA256ToString([]byte("hello world")); s != "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9" {
		t.Errorf("SHA256ToString(hello world) want return b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9, but return %s", s)
	}
}

func BenchmarkMD5ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MD5ToString([]byte("hello world"))
	}
}

func BenchmarkSHA1ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA1ToString([]byte("hello world"))
	}
}

func BenchmarkSHA256ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA256ToString([]byte("hello world"))
	}
}
