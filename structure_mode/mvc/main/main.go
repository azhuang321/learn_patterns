package main

import "go-patterns/structure_mode/mvc"

func main() {
	studentModel := new(mvc.Student)
	studentModel.SetName("Robert")
	studentModel.SetRollNo("10")

	studentView := new(mvc.StudentView)

	studentController := mvc.NewStudent(studentModel, studentView)
	studentController.UpdateView()

	studentController.SetStudentName("john")
	studentController.UpdateView()
}
