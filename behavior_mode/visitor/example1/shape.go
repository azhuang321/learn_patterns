package example1

type Shape interface {
	getType() string
	accept(visitor)
}
