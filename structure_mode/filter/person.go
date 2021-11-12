package filter

type Person struct {
	name, gender, maritalStatus string
}

func NewPerson(name, gender, maritalStatus string) *Person {
	person := &Person{}
	person.name = name
	person.gender = gender
	person.maritalStatus = maritalStatus
	return person
}

func (p *Person) GetName() string {
	return p.name
}
func (p *Person) GetGender() string {
	return p.gender
}
func (p *Person) GetMaritalStatus() string {
	return p.maritalStatus
}
