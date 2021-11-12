package example1

type FactoryProduce struct{}

func (f *FactoryProduce) GetFactory(choice string) AbstractFactory {
	switch choice {
	case "shape":
		return new(ShapeFactory)
	case "color":
		return new(ColorFactory)
	default:
		return nil
	}
}
