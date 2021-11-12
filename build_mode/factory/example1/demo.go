package example1

func main() {
	shapeFactory := &ShapeFactory{}
	circleShape := shapeFactory.GetShape("circle")
	circleShape.Draw()

	squareShape := shapeFactory.GetShape("square")
	squareShape.Draw()
}
