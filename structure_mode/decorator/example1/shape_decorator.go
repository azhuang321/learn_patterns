package example1

import "go-patterns/structure_mode/decorator/shape"

type ShapeDecorator struct {
	DecoratedShape shape.Shaper
}

func (s *ShapeDecorator) ShapeDecorator(decoratedShape shape.Shaper) {
	s.DecoratedShape = decoratedShape
}

func (s *ShapeDecorator) Draw() {
	s.DecoratedShape.Draw()
}
