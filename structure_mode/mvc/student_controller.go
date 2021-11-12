package mvc

type StudentController struct {
	model *Student
	view  *StudentView
}

func NewStudent(model *Student, view *StudentView) *StudentController {
	studentController := &StudentController{}
	studentController.model = model
	studentController.view = view
	return studentController
}

func (s *StudentController) SetStudentName(name string) {
	s.model.SetName(name)
}

func (s *StudentController) GetStudentName() string {
	return s.model.GetName()
}

func (s *StudentController) SetStudentRollNo(rollNo string) {
	s.model.SetRollNo(rollNo)
}

func (s *StudentController) GetStudentRollNo() string {
	return s.model.GetRollNo()
}

func (s *StudentController) UpdateView() {
	s.view.PrintStudentDetail(s.model.GetName(), s.model.GetRollNo())
}
