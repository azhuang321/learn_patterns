package shape

import "fmt"

type ShapeFactory struct {
	circle map[string]*Circle
}

func (s *ShapeFactory) GetCircle(color string) *Circle {
	if s.circle == nil {
		s.circle = make(map[string]*Circle)
	}
	circle, ok := s.circle[color]
	if !ok || circle == nil {
		fmt.Println("creat circle:", color)
		circle = New(color)
		s.circle[color] = circle
	}
	return circle
}
