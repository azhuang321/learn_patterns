package example1

var (
	instance1 *singleton1
)

type singleton1 struct{}

// New1 懒汉模式(lazy loading)	最大的缺点是非线程安全的
func New1() *singleton1 {
	if instance1 == nil {
		instance1 = new(singleton1)
	}
	return instance1
}
