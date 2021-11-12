package shape

import "fmt"

type Circle struct{}

func (r *Circle) Draw() {
	fmt.Println("implements")
}
