package main

import "go-patterns/structure_mode/filter"

func main() {
	persons := make([]*filter.Person, 0)

	persons = append(persons, filter.NewPerson("Robert", "Male", "Single"))
	persons = append(persons, filter.NewPerson("John", "Male", "Married"))
	persons = append(persons, filter.NewPerson("Laura", "Female", "Married"))
	persons = append(persons, filter.NewPerson("Diana", "Female", "Single"))
	persons = append(persons, filter.NewPerson("Mike", "Male", "Single"))
	persons = append(persons, filter.NewPerson("Bobby", "Male", "Single"))

	male := new(filter.CriteriaMale)
	female := new(filter.CriteriaFeMale)
	single := new(filter.CriteriaSingle)

	male.MeetCriteria(persons)
	female.MeetCriteria(persons)
	single.MeetCriteria(persons)
}
