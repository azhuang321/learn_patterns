package example1

type ShapeFactory struct{}

func (s *ShapeFactory) GetShape(shapeName string) Shaper {
	switch shapeName {
	case "rectangle":
		return new(Rectangle)
	case "square":
		return &Square{}
	case "circle":
		return new(Circle)
	default:
		return nil
	}
}
