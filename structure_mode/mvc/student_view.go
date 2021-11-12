package mvc

import "fmt"

type StudentView struct {
}

func (s *StudentView) PrintStudentDetail(studentName string, studentRollNo string) {
	fmt.Println("student:")
	fmt.Println("Name:", studentName)
	fmt.Println("Roll No :", studentRollNo)
}
