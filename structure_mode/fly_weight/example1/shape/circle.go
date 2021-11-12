package shape

import "fmt"

type Circle struct {
	color  string
	x      int
	y      int
	radius int
}

func New(color string) *Circle {
	circle := &Circle{}
	circle.color = color
	return circle
}

func (c *Circle) SetX(x int) {
	c.x = x
}
func (c *Circle) SetY(y int) {
	c.y = y
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) Draw() {
	fmt.Printf("draw x:%d,y:%d,radius:%d\n", c.x, c.y, c.radius)
}
