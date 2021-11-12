package example1

type CareTaker struct {
	mementoList []*Memento
}

func (c *CareTaker) Add(memento *Memento) {
	c.mementoList = append(c.mementoList, memento)
}

func (c *CareTaker) Get(index int) *Memento {
	return c.mementoList[index]
}
