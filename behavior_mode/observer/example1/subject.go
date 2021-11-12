package example1

type Subject struct {
	state    int
	observer []Observer
}

func (s *Subject) GetState() int {
	return s.state
}

func (s *Subject) SetState(state int) {
	s.state = state
	s.NotifyAllObservers()
}

func (s *Subject) Attach(observer Observer) {
	s.observer = append(s.observer, observer)
}

func (s *Subject) NotifyAllObservers() {
	for _, v := range s.observer {
		v.Update()
	}
}
