package mvc

type Student struct {
	rollNo string
	name   string
}

func (s *Student) GetRollNo() string {
	return s.rollNo
}

func (s *Student) SetRollNo(no string) {
	s.rollNo = no
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}
