package example1

import "sync"

var (
	once5     sync.Once
	instance5 *singleton5
)

type singleton5 struct{}

func New5() *singleton5 {
	once5.Do(func() {
		instance5 = &singleton5{}
	})
	return instance5
}
