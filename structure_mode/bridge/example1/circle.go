package example1

type Circle struct {
	Shape
	x, y, radius int
}

func New(x, y, radius int, api DrawAPI) *Circle {
	circle := &Circle{}
	circle.DrawAPI = api
	circle.x = x
	circle.y = y
	circle.radius = radius

	return circle
}

func (c *Circle) Draw() {
	c.DrawAPI.DrawCircle(c.x, c.y, c.radius)
}
