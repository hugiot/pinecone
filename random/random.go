package random

import (
	"math/rand"
	"sync"
	"time"
)

var r *rand.Rand
var l *sync.Mutex

type integer interface {
	int8 | int16 | int | int32 | int64
}

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	l = &sync.Mutex{}
}

func number[T integer](min, max T) T {
	if max < min {
		panic("max must be greater than min")
	}

	l.Lock()
	defer l.Unlock()

	return T(r.Int63n(int64(max)-int64(min)+1) + int64(min))
}

func Int(min, max int) int {
	return number(min, max)
}

func Int64(min, max int64) int64 {
	return number(min, max)
}
