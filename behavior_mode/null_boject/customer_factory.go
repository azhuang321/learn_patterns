package null_boject

import "strings"

type CustomFactory struct {
	name []string
}

func (c *CustomFactory) GetCustomer(name string) AbstractCustomer {
	c.name = []string{"123", "456", "789"}
	for _, v := range c.name {
		if strings.Contains(v, name) {
			return NewRealCustomer(name)
		}
	}
	return NewNullCustomer()
}
