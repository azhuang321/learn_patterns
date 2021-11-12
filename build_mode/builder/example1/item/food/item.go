package food

import (
	"go-patterns/build_mode/builder/pack"
	"go-patterns/build_mode/builder/pack/mode"
)

type Burger struct{}

func (b *Burger) Name() string {
	return "burger"
}

func (b *Burger) Packing() pack.Packinger {
	return new(mode.Wrapper)
}
func (b *Burger) Price() float64 {
	return 0
}

type VegBurger struct {
	Burger
}

func (v *VegBurger) Price() float64 {
	return 26.0
}

func (v *VegBurger) Name() string {
	return "Veg Burger"
}
