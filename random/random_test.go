package random

import "testing"

func TestInt(t *testing.T) {
	for i := 0; i < 10000; i++ {
		num := Int(100, 101)
		if !(num == 100 || num == 101) {
			t.Errorf("Int(100, 101) want return 100 or 101, but return %d", num)
		}
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < 10000; i++ {
		num := Int64(999, 1000)
		if !(num == 999 || num == 1000) {
			t.Errorf("Int(999, 1000) want return 999 or 1000, but return %d", num)
		}
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(100, 200)
	}
}

func BenchmarkInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64(100, 200)
	}
}
