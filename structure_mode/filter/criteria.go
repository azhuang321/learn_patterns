package filter

type Criteria interface {
	MeetCriteria(persons []*Person) []*Person
}

type CriteriaMale struct{}

func (c *CriteriaMale) MeetCriteria(persons []*Person) []*Person {
	malePersons := make([]*Person, 0)
	for _, person := range persons {
		if person.GetGender() == "MALE" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}

type CriteriaFeMale struct{}

func (c *CriteriaFeMale) MeetCriteria(persons []*Person) []*Person {
	malePersons := make([]*Person, 0)
	for _, person := range persons {
		if person.GetGender() == "FEMALE" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}

type CriteriaSingle struct{}

func (c *CriteriaSingle) MeetCriteria(persons []*Person) []*Person {
	malePersons := make([]*Person, 0)
	for _, person := range persons {
		if person.GetGender() == "MALE" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}

type AndCriteria struct{}

func (c *AndCriteria) MeetCriteria(persons []*Person) []*Person {
	malePersons := make([]*Person, 0)
	for _, person := range persons {
		if person.GetGender() == "MALE" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}

type OrCriteria struct{}

func (c *OrCriteria) MeetCriteria(persons []*Person) []*Person {
	malePersons := make([]*Person, 0)
	for _, person := range persons {
		if person.GetGender() == "MALE" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}
