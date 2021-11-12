package example1

import "fmt"

type RedCircle struct {
}

func (r *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Println("draw red circle")
}
