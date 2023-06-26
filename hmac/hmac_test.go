package hmac

import "testing"

func TestSHA1ToString(t *testing.T) {
	if s := SHA1ToString([]byte("123456"), []byte("hello world")); s != "4873e80b320876fef98bd5857a7146db412007c7" {
		t.Errorf("SHA1ToString(123456, hello world) want return 4873e80b320876fef98bd5857a7146db412007c7, but return %s", s)
	}
}

func TestSHA224ToString(t *testing.T) {
	if s := SHA224ToString([]byte("123456"), []byte("hello world")); s != "4f358a44c7bfaf92a2aa3e015d06a6913cc22dd1d684efbb3ce71755" {
		t.Errorf("SHA224ToString(123456, hello world) want return 4f358a44c7bfaf92a2aa3e015d06a6913cc22dd1d684efbb3ce71755, but return %s", s)
	}
}

func TestSHA256ToString(t *testing.T) {
	if s := SHA256ToString([]byte("123456"), []byte("hello world")); s != "83b3eb2788457b46a2f17aaa048f795af0d9dabb8e5924dd2fc0ea682d929fe5" {
		t.Errorf("SHA256ToString(123456, hello world) want return 83b3eb2788457b46a2f17aaa048f795af0d9dabb8e5924dd2fc0ea682d929fe5, but return %s", s)
	}
}

func TestSHA512ToString(t *testing.T) {
	if s := SHA512ToString([]byte("123456"), []byte("hello world")); s != "6c9c251365f3507dc923023fd8e180925eee0dc0bb467d156edc21b9889fc1115cbd7a948090abb59b31718e83978900d7582993392d90d2835ee13c9f2fbb69" {
		t.Errorf("SHA512ToString(123456, hello world) want return 6c9c251365f3507dc923023fd8e180925eee0dc0bb467d156edc21b9889fc1115cbd7a948090abb59b31718e83978900d7582993392d90d2835ee13c9f2fbb69, but return %s", s)
	}
}

func BenchmarkSHA1ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA1ToString([]byte("123456"), []byte("hello world"))
	}
}

func BenchmarkSHA224ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA224ToString([]byte("123456"), []byte("hello world"))
	}
}

func BenchmarkSHA256ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA256ToString([]byte("123456"), []byte("hello world"))
	}
}

func BenchmarkSHA512ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA512ToString([]byte("123456"), []byte("hello world"))
	}
}
