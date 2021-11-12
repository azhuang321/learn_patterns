package example1

import (
	"go-patterns/build_mode/abstract_factory/color"
	"go-patterns/build_mode/abstract_factory/shape"
)

type AbstractFactory interface {
	GetColor(colorName string) color.Colorer
	GetShape(shapeName string) shape.Shaper
}

type ShapeFactory struct{}

func (s *ShapeFactory) GetColor(colorName string) color.Colorer {
	return nil
}

func (s *ShapeFactory) GetShape(shapeName string) shape.Shaper {
	switch shapeName {
	case "rectangle":
		return new(shape.Rectangle)
	case "square":
		return &shape.Square{}
	case "circle":
		return new(shape.Circle)
	default:
		return nil
	}
}

type ColorFactory struct{}

func (s *ColorFactory) GetColor(colorName string) color.Colorer {
	switch colorName {
	case "red":
		return new(color.Red)
	case "green":
		return &color.Green{}
	case "blue":
		return new(color.Blue)
	default:
		return nil
	}
}

func (s *ColorFactory) GetShape(shapeName string) shape.Shaper {
	return nil
}
