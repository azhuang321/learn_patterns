package example1

import "fmt"

type Employee struct {
	name         string
	dept         string
	salary       int
	subordinates []*Employee
}

func New(name, dept string, sal int) *Employee {
	return &Employee{
		name:         name,
		dept:         dept,
		salary:       sal,
		subordinates: []*Employee{},
	}
}

func (e *Employee) Add(ep *Employee) {
	e.subordinates = append(e.subordinates, ep)
}

func (e *Employee) Remove(ep *Employee) {
	for i, v := range e.subordinates {
		if v.name == ep.name {
			e.subordinates = append(e.subordinates[0:i], ep.subordinates[i+1:]...)
			break
		}
	}
}

func (e *Employee) List() {
	fmt.Printf("%+v", e.subordinates)
}
