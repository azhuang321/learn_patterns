package example1

import (
	"sync"
	"sync/atomic"
)

var (
	instance4    *singleton4
	mu4          *sync.Mutex
	initialized4 *uint32
)

type singleton4 struct{}

// New4 原子锁
func New4() *singleton4 {
	if atomic.LoadUint32(initialized4) == 1 {
		return instance4
	}

	mu4.Lock()
	defer mu4.Unlock()

	if *initialized4 == 0 {
		instance4 = new(singleton4)
		atomic.AddUint32(initialized4, 1)
	}
	return instance4
}
