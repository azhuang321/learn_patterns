package example1

import "fmt"

type GreenCircle struct {
}

func (r *GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Println("draw red circle")
}
