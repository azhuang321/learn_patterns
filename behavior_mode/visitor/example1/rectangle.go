package example1

type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
	return "Rectangle"
}
