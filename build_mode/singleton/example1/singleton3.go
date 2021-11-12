package example1

import "sync"

var (
	instance3 *singleton3
	mu3       *sync.Mutex
)

type singleton3 struct{}

// New3 带锁 带检查
func New3() *singleton3 {
	if instance3 == nil {
		mu3.Lock()
		defer mu3.Unlock()

		if instance3 == nil {
			instance3 = new(singleton3)
		}
	}
	return instance3
}
