package example1

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SaveStateToMemento() *Memento {
	return New(o.state)
}

func (o *Originator) GetStateFromMemento(memento *Memento) {
	o.state = memento.state
}
