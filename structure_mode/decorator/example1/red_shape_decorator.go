package example1

import (
	"fmt"
	"go-patterns/structure_mode/decorator/shape"
)

type RedShapeDecorator struct {
	ShapeDecorator
}

func (r *RedShapeDecorator) RedShapeDecorator(decoratedShape shape.Shaper) {
	r.DecoratedShape = decoratedShape
}

func (r *RedShapeDecorator) setRedBorder(decoratedShape shape.Shaper) {
	fmt.Println("border color:red")
}

func (r *RedShapeDecorator) Draw() {
	r.DecoratedShape.Draw()
	r.setRedBorder(r.DecoratedShape)
}
