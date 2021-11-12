package example1

type Memento struct {
	state string
}

func New(state string) *Memento {
	memento := &Memento{}
	memento.state = state
	return memento
}

func (m *Memento) GetState() string {
	return m.state
}
