package example1

import (
	"go-patterns/structure_mode/facade/shape"
)

type ShapeMaker struct {
	circle    shape.Shaper
	square    shape.Shaper
	rectangle shape.Shaper
}

func New() *ShapeMaker {
	circle := new(shape.Circle)
	square := new(shape.Square)
	rectangle := new(shape.Rectangle)

	shapeMaker := &ShapeMaker{}
	shapeMaker.square = square
	shapeMaker.circle = circle
	shapeMaker.rectangle = rectangle
	return shapeMaker
}

func (s *ShapeMaker) DrawCircle() {
	s.circle.Draw()
}

func (s *ShapeMaker) DrawSquare() {
	s.square.Draw()
}

func (s *ShapeMaker) DrawRectangle() {
	s.rectangle.Draw()
}
