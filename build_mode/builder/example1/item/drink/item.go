package drink

import (
	"go-patterns/build_mode/builder/pack"
	"go-patterns/build_mode/builder/pack/mode"
)

type ColdDrink struct{}

func (b *ColdDrink) Name() string {
	return "burger"
}

func (b *ColdDrink) Packing() pack.Packinger {
	return new(mode.Wrapper)
}
func (b *ColdDrink) Price() float64 {
	return 0
}

type Coke struct {
	ColdDrink
}

func (c *Coke) Price() float64 {
	return 60
}

func (c *Coke) Name() string {
	return "coke"
}
