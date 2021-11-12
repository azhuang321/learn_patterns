package main

import "go-patterns/structure_mode/facade"

func main() {
	shapeMaker := facade.New()
	shapeMaker.DrawCircle()
	shapeMaker.DrawRectangle()
	shapeMaker.DrawSquare()
}
