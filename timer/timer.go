package timer

import (
	"context"
	"sync"
	"time"
)

var counter int64
var list = make(map[int64]context.CancelFunc)
var lock sync.Mutex

func SetInterval(f func(), t time.Duration) int64 {
	lock.Lock()
	defer lock.Unlock()

	counter++
	ctx, cancel := context.WithCancel(context.Background())
	list[counter] = cancel

	ticker := time.NewTicker(t)
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			case <-ctx.Done():
				goto overIntervalTag
			}
		}
	overIntervalTag:
		ticker.Stop()
	}()

	return counter
}

func SetTimeout(f func(), t time.Duration) int64 {
	lock.Lock()
	defer lock.Unlock()

	counter++
	ctx, cancel := context.WithCancel(context.Background())
	list[counter] = cancel

	timer := time.NewTimer(t)
	go func() {
		select {
		case <-timer.C:
			f()
		case <-ctx.Done():
			// nothing todo
		}
		timer.Stop()
	}()

	return counter
}

func Clear(id int64) {
	lock.Lock()
	defer lock.Unlock()

	if cancel, ok := list[id]; ok {
		cancel()
		delete(list, id)
	}
}
