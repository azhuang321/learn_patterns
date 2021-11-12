package main

import (
	"fmt"
	. "go-patterns/structure_mode/decorator"
	"go-patterns/structure_mode/decorator/shape"
)

func main() {
	circle := new(shape.Circle)
	redCircle := &RedShapeDecorator{ShapeDecorator: ShapeDecorator{DecoratedShape: new(shape.Circle)}}
	redRectangle := &RedShapeDecorator{ShapeDecorator: ShapeDecorator{DecoratedShape: new(shape.Rectangle)}}

	fmt.Println("normal circle")
	circle.Draw()
	fmt.Println()

	fmt.Println("red circle")
	redCircle.Draw()
	fmt.Println()

	fmt.Println("red rectangle")
	redRectangle.Draw()
	fmt.Println()
}
