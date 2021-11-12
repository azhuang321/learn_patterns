package color

import "fmt"

type Red struct{}

func (r *Red) Fill() {
	fmt.Println("implements")
}
