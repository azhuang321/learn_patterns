package example1

type Example struct {
	Content string
}

func (e *Example) Clone() *Example {
	res := *e
	return &res
}
