package example1

import "sync"

var (
	instance2 *singleton2
	mu2       *sync.Mutex
)

type singleton2 struct{}

// New2 带锁
func New2() *singleton2 {
	mu2.Lock()
	defer mu2.Unlock()

	if instance2 == nil {
		instance2 = new(singleton2)
	}
	return instance2
}
