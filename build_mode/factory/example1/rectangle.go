package example1

import "fmt"

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("implements")
}
