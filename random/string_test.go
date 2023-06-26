package random

import "testing"

func TestString(t *testing.T) {
	if s := String(10); len(s) != 10 {
		t.Errorf("String(10) want return len is 10, but return %d", len(s))
	}

	if s := String(888); len(s) != 888 {
		t.Errorf("String(10) want return len is 888, but return %d", len(s))
	}

	if s := String(9999); len(s) != 9999 {
		t.Errorf("String(10) want return len is 9999, but return %d", len(s))
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(999)
	}
}
