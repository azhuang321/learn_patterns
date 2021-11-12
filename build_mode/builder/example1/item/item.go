package item

import "go-patterns/build_mode/builder/pack"

type Itemer interface {
	Name() string
	Packing() pack.Packinger
	Price() float64
}
